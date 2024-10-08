from django.db import models
from django.contrib.auth.models import User

class Problema(models.Model):
    id = models.AutoField(primary_key=True)
    enunciado = models.TextField()
    solucion = models.TextField()
    titulo = models.CharField(max_length=255)
    fecha_desbloqueo = models.DateTimeField()

    def __str__(self):
        return self.titulo
    
class Respuesta(models.Model):
    id = models.AutoField(primary_key=True)
    usuario = models.ForeignKey(User, on_delete=models.CASCADE)
    problema = models.ForeignKey(Problema, on_delete=models.CASCADE)

    solucion_propuesta = models.TextField()
    fecha_envio = models.DateTimeField(auto_now_add=True)
    correcta = models.BooleanField(default=False)

    def __str__(self):
        return f"{self.usuario.username} - {self.problema.titulo}"

