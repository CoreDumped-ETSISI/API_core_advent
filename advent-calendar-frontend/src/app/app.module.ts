// src/app/app.module.ts
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './components/app/app.component';
import { LoginComponent } from './components/login/login.component';
import { FormBuilder, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { provideHttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';

@NgModule({

  declarations: [
    AppComponent, 
    LoginComponent,
    ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [provideHttpClient()],
  bootstrap: [AppComponent],
})
export class AppModule {}
