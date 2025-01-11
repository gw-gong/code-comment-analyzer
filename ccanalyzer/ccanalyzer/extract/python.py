# encoding=utf-8
# 创建时间：2024-04-29 11:46

# ---------------- | 功能 | --------------------
# 提取 python 的注释
# ---------------------------------------------
import re

from .base import CommentExtractor
import ast
from .result import CommentExtractorResult


class PythonCommentExtractor(CommentExtractor):
    def __init__(self):
        super().__init__()
        # self.single_line_pattern = r'#\s*(.*?)(?=\n|$)'
        self._single_line_pattern = r'\#(.*?)\n'
        self._multi_line_pattern1 = r'"""([\s\S]*?)"""'
        self._multi_line_pattern2 = r'\'\'\'([\s\S]*?)\'\'\''

    """ 提取Python代码中的字符串 """
    def extract_strings(self, source_code):
        tree = ast.parse(source_code)
        strings = []

        class StringVisitor(ast.NodeVisitor):
            def visit_Expr(self, node):
                # 仅当Expr内部是一个字符串时，考虑它可能是注释
                if isinstance(node.value, (ast.Str, ast.Constant)) and isinstance(node.value.value, str):
                    if '\n' in node.value.value:
                        # 可能是多行注释，不添加
                        return
                self.generic_visit(node)

            def visit_Str(self, node):
                # 处理Python 3.7及之前版本的字符串节点
                if node.s.strip():
                    strings.append((node.s, node.lineno, node.col_offset))

            def visit_Constant(self, node):
                # 处理Python 3.8及之后版本的字符串常量节点
                if isinstance(node.value, str) and node.value.strip():
                    strings.append((node.value, node.lineno, node.col_offset))

        StringVisitor().visit(tree)
        return strings

    """ 从源码字符串中去除字符串 """
    def remove_strings_from_code(self, source_code, strings_info):
        """
        Remove strings from source code by replacing them with spaces.
        Assumes multiline strings are included with their content intact including newline characters.
        :param source_code: The original source code.
        :param strings_info: A list of tuples containing string content, line number, and start column.
        :return: Modified source code with strings replaced by spaces.
        """
        # 将源代码按行分割成列表，便于处理
        lines = source_code.split('\n')

        # 从后往前替换字符串，以免修改行的长度影响后续字符串的位置
        for string_content, line_no, start_col in reversed(strings_info):
            # 修正一下位置, 和字符串
            line = lines[line_no - 1]
            # print(string_content, start_col)
            cnt = 0
            if line[start_col] == '"':
                for i in range(0, 3):
                    if line[start_col] == '"':
                        string_content = '"' + string_content + '"'
                        cnt += 1
                        start_col += 1
                    else:
                        break

            elif line[start_col] == "'":
                for i in range(0, 3):
                    if line[start_col] == "'":
                        string_content = "'" + string_content + "'"
                        cnt += 1
                        start_col += 1
                    else:
                        break
            start_col -= cnt
            # print(string_content, cnt)

            # 计算字符串跨越的行数
            string_lines = string_content.split('\n')
            start_line_index = line_no - 1
            end_line_index = start_line_index + len(string_lines) - 1

            # 处理跨越多行的字符串
            for i in range(len(string_lines)):
                current_line_index = start_line_index + i
                if i == 0:
                    # 首行
                    current_line_part = string_lines[i]
                    end_col = start_col + len(current_line_part)
                    lines[current_line_index] = (lines[current_line_index][:start_col] +
                                                '' * len(current_line_part) +
                                                lines[current_line_index][end_col:])

                else:
                    # 中间行到结束行
                    current_line_part = string_lines[i]
                    if i == len(string_lines) - 1:
                        # 最后一行
                        lines[current_line_index] = ('' * len(current_line_part) +
                                                    lines[current_line_index][len(current_line_part):])
                    else:
                        # 完整行
                        lines[current_line_index] = '' * len(lines[current_line_index])

        # 将修改后的行重新组合成完整的代码文本
        modified_code = '\n'.join(lines)
        return modified_code

    """" 提取注释 """

    def extract_comments(self, code):
        result = CommentExtractorResult()
        result.setCodeLines(code.count('\n') + 1)  # 代码行数，加一是因为最后一行没有换行符

        # 提取代码中的字符串
        strings = self.extract_strings(code)

        # print(strings)

        # 从源码中删除字符串
        code = self.remove_strings_from_code(code, strings)

        # print(code)

        pattern = ""
        if self._single_line_pattern is not None:
            pattern += self._single_line_pattern
        if self._multi_line_pattern1 is not None:
            pattern += '|'
            pattern += self._multi_line_pattern1
        if self._multi_line_pattern2 is not None:
            pattern += '|'
            pattern += self._multi_line_pattern2

        matches = re.finditer(pattern, code)

        SingleLineComments = []
        MultiLineComments1 = []
        MultiLineComments2 = []

        for match in matches:
            single_line, double_quotes, single_quotes = match.groups()
            if single_line:
                # print(f"Single-line comment: {single_line}")
                SingleLineComments.append(single_line)

            elif double_quotes:
                # print(f"Double-quoted multi-line comment: {double_quotes}")
                MultiLineComments1.append(double_quotes)
            elif single_quotes:
                # print(f"Single-quoted multi-line comment: {single_quotes}")
                MultiLineComments2.append(single_quotes)

        result.setSingleLineComments(SingleLineComments)
        result.setMultiLineComments1(MultiLineComments1)
        result.setMultiLineComments2(MultiLineComments2)

        return result
