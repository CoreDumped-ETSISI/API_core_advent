// src/store.js
import { writable } from 'svelte/store';

export const isAdmin = writable(false); // Default to not being an admin

export const yearStore = writable(null);
