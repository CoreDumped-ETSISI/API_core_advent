// src/app/services/api.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl = 'http://localhost:8000'; 

  constructor(private http: HttpClient) {}

  login(usuario: string, contrasena: string): Observable<any> {

    console.log(usuario, contrasena);
    return this.http.post(`${this.apiUrl}/login`, 
      { Valor: usuario,
        Contraseña: contrasena 
      },
      { responseType: 'json' },);
    }

  register(correo:string, contrasena: string,  usuario: string): Observable<any> {
    // imprime los valores de correo, contrasena y usuario
    console.log(correo, contrasena, usuario);
    
    return this.http.post(`${this.apiUrl}/register`, 
      { Correo: correo,
        Usuario: usuario,
        Constraseña: contrasena 
      },
      { responseType: 'json' },

    );
    }

  // Agrega otros métodos para manejar problemas y respuestas
}
