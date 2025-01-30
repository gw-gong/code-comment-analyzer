# encoding=utf-8
# 创建时间：2024-04-29 11:44

# ---------------- | 功能 | --------------------
# 提取 C 或 C++ 的注释
# ---------------------------------------------
import re

from .base import CommentExtractor
from .result import CommentExtractorResult


class CCppCommentExtractor(CommentExtractor):
    def __init__(self):
        super().__init__()
        # 单行注释匹配从 `//` 开始到行末的所有内容
        self._single_line_pattern = r'\/\/(.*?)\n'
        # 多行注释匹配从 `/*` 开始到 `*/` 结束的所有内容，包括换行符
        self._multi_line_pattern1 = r'\/\*([\s\S]*?)\*\/'
        # 字符串规则
        self.__string_pattern = r'\"([\s\S]*?)\"'

    """" 提取注释 """
    def extract_comments(self, code):
        result = CommentExtractorResult()
        result.setCodeLines(code.count('\n') + 1)  # 代码行数，加一是因为最后一行没有换行符

        pattern = ""

        if self._single_line_pattern is not None:
            pattern += self._single_line_pattern
        if self._multi_line_pattern1 is not None:
            pattern += '|'
            pattern += self._multi_line_pattern1
        if self.__string_pattern is not None:
            pattern += '|'
            pattern += self.__string_pattern

        matches = re.finditer(pattern, code)

        SingleLineComments = []
        MultiLineComments1 = []

        for match in matches:
            single_line, mult_line, str_line = match.groups()
            if single_line:
                # print(f"Single-line comment: {single_line}")
                SingleLineComments.append(single_line)
            elif mult_line:
                # print(f"Double-quoted multi-line comment: {double_quotes}")
                MultiLineComments1.append(mult_line)

        result.setSingleLineComments(SingleLineComments)
        result.setMultiLineComments1(MultiLineComments1)

        return result






