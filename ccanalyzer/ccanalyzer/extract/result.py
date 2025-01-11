# encoding=utf-8
# 创建时间：2024-04-30 14:50

# ---------------- | 功能 | --------------------
# 处理返回的结果 CommentExtractorResult
# ---------------------------------------------

class CommentExtractorResult:
    def __init__(self):
        self.__codeLines = 0                # 代码行数
        self.__single_line_comments = []    # 单行注释
        self.__multi_line_comments1 = []    # 多行注释1
        self.__multi_line_comments2 = []    # 多行注释2

    def setCodeLines(self, codeLines):
        self.__codeLines = codeLines

    def setSingleLineComments(self, SingleLineComments):
        self.__single_line_comments = SingleLineComments

    def setMultiLineComments1(self, multi_line_comments):
        self.__multi_line_comments1 = multi_line_comments

    def setMultiLineComments2(self, multi_line_comments):
        self.__multi_line_comments2 = multi_line_comments

    def getCodeLines(self):
        return self.__codeLines

    def getSingleLineComments(self):
        return self.__single_line_comments

    def getMultiLineComments(self):
        return self.__multi_line_comments1 + self.__multi_line_comments2

    """ 打印单行注释 """
    def printSingleLineComments(self):
        print('------------------------------')
        print('单行注释：', len(self.__single_line_comments), '个')
        for comment in self.__single_line_comments:
            print(comment)
        print('\n\n\n')

    """ 打印多行注释 """
    def printMultiLineComments(self):
        print('------------------------------')
        print('多行注释：', len(self.__multi_line_comments1) + len(self.__multi_line_comments2), '个')
        for comment in self.__multi_line_comments1:
            print(comment)
        for comment in self.__multi_line_comments2:
            print(comment)
        print('\n\n\n')