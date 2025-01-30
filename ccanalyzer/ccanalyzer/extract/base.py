# encoding=utf-8
# 创建时间：2024-04-29 14:32

# ---------------- | 功能 | --------------------
# 基类 CommentExtractor
# ---------------------------------------------

import re
from .result import CommentExtractorResult

class CommentExtractor:
    def __init__(self):
        self._single_line_pattern = None
        self._multi_line_pattern1 = None     # 多行注释1
        self._multi_line_pattern2 = None     # 多行注释2, 用来兼容其他语言的多行注释





