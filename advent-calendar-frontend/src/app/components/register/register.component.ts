import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ApiService } from '../../services/api.service';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  imports : [CommonModule, ReactiveFormsModule],
  standalone: true,
})
export class RegisterComponent {
  registerForm: FormGroup;

  constructor(private fb: FormBuilder, private api: ApiService, private router: Router) {
    this.registerForm = this.fb.group({
      correo: ['', Validators.required],
      usuario: ['', Validators.required],
      contrasena: ['', Validators.required],
    });
  }

  onSubmit(): void {
    if (this.registerForm.valid) {
      const correo = this.registerForm.get('correo')?.value;
      const usuario = this.registerForm.get('usuario')?.value;
      const contrasena = this.registerForm.get('contrasena')?.value;
      this.api.register(correo, contrasena, usuario).subscribe({
        next: (response) => {
          console.log('Registro exitoso', response);
          // Aquí puedes redirigir al usuario o mostrar un mensaje
          this.router.navigate(['/login']);
        },
        error: (err) => {
          console.error('Error en el registro', err);
        },
      });
    }
  }
}
