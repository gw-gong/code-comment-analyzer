# encoding=utf-8
# 创建时间：2024-04-29 23:05

# ---------------- | 功能 | --------------------
# 分析处理提取的注释
# ---------------------------------------------

import jieba
import re
from nltk.corpus import stopwords  # 导入停用词典
from nltk.tokenize import word_tokenize
from collections import Counter


class Processor:

    def __init__(self, result):
        self._codeLines = result.getCodeLines()
        self._single_line_comments = result.getSingleLineComments()  # 单行注释
        self._multi_line_comments = result.getMultiLineComments()  # 多行注释
        self._commentString = ''  # 评论总体字符串
        self._words = []  # 分词结果

    """ 注释行数 """
    def getCommentLines(self):
        sum = len(self._single_line_comments)
        for comment in self._multi_line_comments:
            # 计算换行符数量
            number_of_lines = comment.count('\n') + 1  # 加1是因为最后一行没有换行符
            sum += number_of_lines
        # print('注释行数', sum, ' 代码行数', self._codeLines)
        return sum

    """ 代码行数 """
    def getCodeLines(self):
        return self._codeLines

    """ 主要语言 """
    # def getMainLanguage(self):
    #     comments = self.__getString()
    #     # 检测语言
    #     # zh-cn
    #     # en
    #     # print('---------------------')
    #     # print(comments)
    #     # DetectorFactory.seed = 0  # 设置固定种子以减少随机性
    #     detector = Detector(comments)
    #     return detector.language.code

    """ 去停用词前 关键字提取 , 最多n个关键字"""
    def getKeyWordsBefore(self, n):
        if len(self._words) == 0:
            self.__dividWords()
        word_counts = Counter(self._words)
        top_n_words = word_counts.most_common(n)
        return top_n_words

    """ 去停用词后 关键字提取 , 最多n个关键字"""
    def getKeyWordsAfter(self, n):
        if len(self._words) == 0:
            self.__dividWords()
        words = self.__getRemoveStopwords()

        word_counts = Counter(words)
        top_n_words = word_counts.most_common(n)
        return top_n_words

    """ ----------------------------------------------- """
    """ 返回评论总体字符串 """
    def __getString(self):
        if self._commentString != '':
            return self._commentString

        for comment in self._single_line_comments:
            self._commentString += comment
        for comment in self._multi_line_comments:
            self._commentString += comment
        return self._commentString

    """ 分词 """
    def __dividWords(self):
        if len(self._words) != 0:
            return self._words

        # 用于检查一个词是否为有效内容（不包括纯标点）
        def is_valid_word(word):
            return re.search('[a-zA-Z0-9\u4e00-\u9fff]', word) is not None

        commentString = self.__getString()
        # 使用jieba分词进行中文分词
        self._words = list(jieba.cut(commentString))
        # 使用nltk分词进行英文分词，并过滤标点
        self._words = [word for segment in self._words for word in word_tokenize(segment) if is_valid_word(word)]

        return self._words

    """ 去停用词 """

    def __getRemoveStopwords(self):
        if len(self._words) == 0:
            self.__dividWords()

        # 加载中文和英文停用词
        stop_words = set(stopwords.words('chinese')).union(set(stopwords.words('english')))

        # 过滤停用词
        filtered_words = [word for word in self._words if word not in stop_words]
        return filtered_words