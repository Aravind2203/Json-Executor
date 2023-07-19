from django.shortcuts import render
from .utils import *
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
import json
# Create your views here.


class ExecuteJson(APIView):
    authentication_classes=[]
    permission_classes=[]

    def get(self,request):

        return Response(data={
            "message":"This is the data"
        },status=status.HTTP_200_OK)

    def post(self,request):
        data=self.request.data
        rpc=RabbitMqConnection()
        rpc.call(json.dumps(data))
        resp=rpc.response
        return Response(json.loads(json.loads(resp)) ,status=status.HTTP_200_OK)