import grpc
from concurrent import futures
from grpc_server import ccanalyzer_pb2
from grpc_server import ccanalyzer_pb2_grpc

from ccanalyzer import extract
from ccanalyzer import process

from config import config
import json

import nltk
import ssl
try:
    _create_unverified_https_context = ssl._create_unverified_context
except AttributeError:
    pass
else:
    ssl._create_default_https_context = _create_unverified_https_context
nltk.download('punkt')  # 确保punkt资源已经下载
nltk.download('stopwords')  # 确保stopwords资源已经下载


class CcAnalyzer(ccanalyzer_pb2_grpc.CcAnalyzerServicer):
    def __init__(self, maxNum):
        super().__init__()
        self.__maxNum = maxNum
    def AnalyzeFileContent(self, request: ccanalyzer_pb2.AnalyzeFileContentReq, context):
        language = request.language
        file_content = request.file_content
        if language == "Unknown":
            return ccanalyzer_pb2.AnalyzeFileContentRes(analyzed_data="")

        codeExtract = extract.PythonCommentExtractor()
        if language == "Python":
            codeExtract = extract.PythonCommentExtractor()
        if language == "C":
            codeExtract = extract.CCppCommentExtractor()
        if language == "CPP":
            codeExtract = extract.CCppCommentExtractor()
        if language == "Java":
            codeExtract = extract.JavaCommentExtractor()
        if language == "Go":
            codeExtract = extract.GoCommentExtractor()
        if language == "HTML":
            codeExtract = extract.HTMLCommentExtractor()
        if language == "CSS":
            codeExtract = extract.CSSCommentExtractor()
        if language == "JavaScript":
            codeExtract = extract.JavaScriptCommentExtractor()

        ceResult = codeExtract.extract_comments(file_content)

        processedJson = process.Json(ceResult)

        analyzedData = processedJson.getJson(self.__maxNum)

        # 转换换行符号位html符号
        for i in range(len(analyzedData['multiLineComments'])):
            analyzedData['multiLineComments'][i] = analyzedData['multiLineComments'][i].replace("\n", "<br>").replace(" ", "&nbsp;")
        return ccanalyzer_pb2.AnalyzeFileContentRes(analyzed_data=analyzedData)


def serve():
    cfg = config.GetConfig("./config/config.yaml")
    # 使用线程池来完成grpc的请求
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    ccanalyzer_pb2_grpc.add_CcAnalyzerServicer_to_server(CcAnalyzer(cfg["max_num_of_word_frequency"]), server)
    print('listening ' + '127.0.0.1:8888')
    server.add_insecure_port('127.0.0.1:8888')  # 绑定端口
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print('start')
    serve()