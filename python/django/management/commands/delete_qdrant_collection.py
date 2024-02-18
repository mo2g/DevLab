from django.core.management.base import BaseCommand, CommandError
import os
import qdrant_client

QDRANT_HOST = os.environ.get("QDRANT_HOST", 'qdrant')

EMBEDDING_MODEL = {
    'text-embedding-ada-002': 8191,
    'text-embedding-3-large': 8191,
    'text-embedding-3-small': 8191,
    'text-embedding-v2': 2047,
    'embedding-2': 512,
}

class Command(BaseCommand):
    help = '遍历集合，删掉 不以向量模型作为后缀的集合'

    # 添加额外的参数
    def add_arguments(self, parser):
        parser.add_argument('--par', dest='par', default='all')

    # 重写父类的handle方法，继承BaseCommand必须重写该方法，该方法就是执行的代码
    def handle(self, **options):
        # 初始化 Qdrant 客户端
        client = qdrant_client.QdrantClient(
            host=QDRANT_HOST,
            prefer_grpc=False
        )

        collections = client.get_collections()

        embedding_model_names = EMBEDDING_MODEL.keys()

        for collection in collections.collections:
            # 遍历集合，删掉 不以向量模型作为后缀的集合
            collection_name = collection.name
            if not any(embedding_model in collection_name for embedding_model in embedding_model_names):
                client.delete_collection(collection_name=collection_name)
                print(f"Collection '{collection_name}' has been deleted.")

        print("Operation completed.")