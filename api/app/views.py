from urllib import request
from django.shortcuts import render
from rest_framework import generics
from .models import Problema
from .serializers import ProblemaSerializer, RespuestaSerializer

class ProblemView(generics.ListAPIView):

    queryset = Problema.objects.all()
    serializer_class = ProblemaSerializer


class RespuestaView(generics.ListCreateAPIView):

    queryset = Problema.objects.filter(usuario=request.user)
    serializer_class = RespuestaSerializer

