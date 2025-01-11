# encoding=utf-8
# 创建时间：2024-04-30 14:45

# ---------------- | 功能 | --------------------
# 返回需要的数据
# ---------------------------------------------
import re

import jieba
from nltk import word_tokenize

from .processor import Processor


class Json(Processor):

    def __init__(self, result):
        super().__init__(result)
        self.__commentDensity = ""
        # self.__language = ' '
        self.__json = {}

    """
        核心函数，返回json格式的数据
        {
            "commentLines":     注释行数,
            "codeLines":        代码行数,
            "commentDensity":   注释密度,
            "lenWordsBefore":   对应长度,
            "wordsBefore": [    去停用词前 关键字
                {
                    "key": "xxxx",  关键字
                    "value": 666    关键字出现频率
                },
                ……
            ],      
            "lenWordsAfter":    对应长度,
            "wordsAfter":       去停用词后 关键字
            {
                "keyword1":
                ……
            },
            "singleLineCommentsLength": ,     单行注释数
            "singleLineComments": [],           单行注释
            "multiLineCommentsLength": ,      多行注释数 
            "multiLineComments": [],            多行注释 
        }
    """

    def getJson(self, maxNum):
        if len(self.__json) > 0:
            return self.__json

        self.__json['analyzed_table'] = []
        self.__json['analyzed_table'].append({
            "key": "代码注释行数",
            "value": self.getCommentLines(),
            "badgeText": "图",
            "badgeLevel": "info"
        })

        self.__json['analyzed_table'].append({
            "key": "代码行数",
            "value": self._codeLines,
            "badgeText": "图",
            "badgeLevel": "info"
        })

        self.__json['commentLines'] = self.getCommentLines()
        self.__json['codeLines'] = self._codeLines

        if self.__commentDensity == "":
            commentDensity = (self.getCommentLines() / self.getCodeLines()) * 100
            self.__commentDensity = f'{commentDensity:.2f}%'
        self.__json['commentDensity'] = self.__commentDensity
        self.__json['analyzed_table'].append({
            "key": "注释覆盖比",
            "value": self.__commentDensity,
            "badgeText": "重点",
            "badgeLevel": "success"
        })

        keyWordsBefore = self.getKeyWordsBefore(maxNum)
        self.__json['lenWordsBefore'] = len(keyWordsBefore)
        self.__json['wordsBefore'] = []
        for keyword in keyWordsBefore:
            tempDict = {'key': keyword[0], 'value': keyword[1]}
            self.__json['wordsBefore'].append(tempDict)
            # self.__json['wordsBefore'][keyword[0]] = keyword[1]

        keyWordsAfter = self.getKeyWordsAfter(maxNum)
        self.__json['lenWordsAfter'] = len(keyWordsAfter)
        self.__json['wordsAfter'] = []
        for keyword in keyWordsAfter:
            tempDict = {'key': keyword[0], 'value': keyword[1]}
            self.__json['wordsAfter'].append(tempDict)
            # self.__json['wordsAfter'][keyword[0]] = keyword[1]

        self.__json['singleLineCommentsLength'] = len(self._single_line_comments)
        self.__json['singleLineComments'] = self._single_line_comments
        self.__json['multiLineCommentsLength'] = len(self._multi_line_comments)
        self.__json['multiLineComments'] = self._multi_line_comments

        word0_1 = 0
        word2_30 = 0
        word31_ = 0
        maxWordNum = 0

        if len(self._single_line_comments) != 0:
            allSingleLineCommentLength = 0
            for single_line_comment in self._single_line_comments:
                allSingleLineCommentLength += len(single_line_comment)

                # 用于检查一个词是否为有效内容（不包括纯标点）
                def is_valid_word(word):
                    return re.search('[a-zA-Z0-9\u4e00-\u9fff]', word) is not None

                # 使用jieba分词进行中文分词
                words = list(jieba.cut(single_line_comment))
                # 使用nltk分词进行英文分词，并过滤标点
                words = [word for segment in words for word in word_tokenize(segment) if
                         is_valid_word(word)]
                if len(words) < 2:
                    word0_1 += 1
                elif 2 <= len(words) <= 30:
                    word2_30 += 1
                else:
                    word31_ += 1
                maxWordNum = max(maxWordNum, len(words))

            self.__json['analyzed_table'].append({
                "key": "单行注释个数",
                "value": len(self._single_line_comments),
                "badgeText": "图",
                "badgeLevel": "info"
            })
            self.__json['analyzed_table'].append({
                "key": "单行注释平均长度",
                "value": f'{(allSingleLineCommentLength / len(self._single_line_comments)):.2f}个字符'
            })
        else:
            self.__json['analyzed_table'].append({
                "key": "单行注释个数",
                "value": "-",
                "badgeText": "图",
                "badgeLevel": "info"
            })
            self.__json['analyzed_table'].append({
                "key": "单行注释平均长度",
                "value": "-"
            })

        if len(self._multi_line_comments) != 0:
            allMultiLineCommentLength = 0
            for multi_line_comment in self._multi_line_comments:
                # print(multi_line_comment,len(multi_line_comment))
                allMultiLineCommentLength += len(multi_line_comment)

                # 用于检查一个词是否为有效内容（不包括纯标点）
                def is_valid_word(word):
                    return re.search('[a-zA-Z0-9\u4e00-\u9fff]', word) is not None

                # 使用jieba分词进行中文分词
                words = list(jieba.cut(multi_line_comment))
                # 使用nltk分词进行英文分词，并过滤标点
                words = [word for segment in words for word in word_tokenize(segment) if
                         is_valid_word(word)]
                if len(words) < 2:
                    word0_1 += 1
                elif 2 <= len(words) <= 30:
                    word2_30 += 1
                else:
                    word31_ += 1
                maxWordNum = max(maxWordNum, len(words))

            self.__json['analyzed_table'].append({
                "key": "多行注释个数",
                "value": len(self._multi_line_comments),
                "badgeText": "图",
                "badgeLevel": "info"
            })
            self.__json['analyzed_table'].append({
                "key": "多行注释平均长度",
                "value": f'{(allMultiLineCommentLength / len(self._multi_line_comments)):.2f}个字符'
            })
        else:
            self.__json['analyzed_table'].append({
                "key": "多行注释个数",
                "value": "-",
                "badgeText": "图",
                "badgeLevel": "info"
            })
            self.__json['analyzed_table'].append({
                "key": "多行注释平均长度",
                "value": "-"
            })

        self.__json['analyzed_table'].append({
            "key": "词数<2 的注释",
            "value": word0_1,
            "badgeText": "图",
            "badgeLevel": "info"
        })
        self.__json['analyzed_table'].append({
            "key": "2<=词数<=30 的注释",
            "value": word2_30,
            "badgeText": "图",
            "badgeLevel": "info"
        })
        self.__json['analyzed_table'].append({
            "key": "词数>30 的注释",
            "value": word31_,
            "badgeText": "图",
            "badgeLevel": "info"
        })
        self.__json['analyzed_table'].append({
            "key": "注释的最大词数",
            "value": maxWordNum
        })

        return self.__json
