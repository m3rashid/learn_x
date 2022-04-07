from django.urls import path
from . import views



urlpatterns = [
    path('', views.getItem),
    path('add/', views.addItem),
]