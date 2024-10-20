// src/app/pages/login/login.component.ts
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, ReactiveFormsModule } from '@angular/forms';
import { ApiService } from '../../services/api.service';
import { CommonModule } from '@angular/common';
import { Router} from '@angular/router';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  imports : [CommonModule, ReactiveFormsModule],
  standalone: true,

})
export class LoginComponent  {
  loginForm: FormGroup;

  constructor(private fb: FormBuilder, private api: ApiService, private router: Router) {
    this.loginForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  ngOnInit(): void {}

  onSubmit(): void {
    if (this.loginForm.valid) {
      const correo = this.loginForm.get('username')?.value;
      const contrasena = this.loginForm.get('password')?.value;
      console.log('Correo:', correo, 'Contraseña:', contrasena);
      this.api.login(correo, contrasena).subscribe({
        next: (response) => {
          console.log('Login exitoso');
          // Aquí puedes manejar la respuesta y almacenar el token si es necesario
          // GUARDAR EL TOKEN
          localStorage.setItem('token', response.token);
          this.router.navigate(['/home']);
        },
        error: (err) => {
          console.error('Error en el login', err);
        },
      });
    }
  }
}