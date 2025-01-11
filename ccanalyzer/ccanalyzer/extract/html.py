# encoding=utf-8
# 创建时间：2024-04-29 11:46

# ---------------- | 功能 | --------------------
# 提取 html 的注释
# ---------------------------------------------
import re

from .base import CommentExtractor
from .result import CommentExtractorResult


class HTMLCommentExtractor(CommentExtractor):
    def __init__(self):
        super().__init__()
        self._multi_line_pattern1 = r'<!--([\s\S]*?)-->'

    """" 提取注释 """
    def extract_comments(self, code):
        result = CommentExtractorResult()
        result.setCodeLines(code.count('\n') + 1)  # 代码行数，加一是因为最后一行没有换行符

        pattern = ""

        if self._multi_line_pattern1 is not None:
            pattern += self._multi_line_pattern1

        matches = re.finditer(pattern, code)

        MultiLineComments1 = []

        for match in matches:
            mult_line = match.groups()
            if mult_line:
                MultiLineComments1.append(mult_line[0])

        # print(MultiLineComments1)
        result.setMultiLineComments1(MultiLineComments1)

        return result
