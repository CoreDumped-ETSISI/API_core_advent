from APP import views
from django.urls import path

urlpatterns = [
    path("problem"),
    path("problem/<int:id>"),
    path("signup"),
    path("signin"),
    path("submit/<int:id>"),
]