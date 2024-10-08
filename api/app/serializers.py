from .models import Problema, Respuesta
from rest_framework import serializers

class ProblemaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Problema
        fields = "__all__"

class RespuestaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Respuesta
        fields = "__all__"
