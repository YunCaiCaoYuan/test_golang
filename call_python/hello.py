# coding: utf-8

from chatterbot import ChatBot
from chatterbot.trainers import ChatterBotCorpusTrainer

def get_response():
    chatbot = ChatBot(
        'xiaojuzi',
        storage_adapter='chatterbot.storage.SQLStorageAdapter',
        database_uri='sqlite:///xiaojuzi.sqlite3'
        # storage_adapter='chatterbot.storage.MongoDatabaseAdapter',
        # database_uri='mongodb://localhost:27017/xiaojuzi'
    )
    trainer = ChatterBotCorpusTrainer(chatbot)
    trainer.train("chatterbot.corpus.chinese") # Training xxx.yml ...
    res = chatbot.get_response("什么是ai")
    print(res)
    return res
