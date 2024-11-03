// +page.server.ts
import { API_URI } from "$env/static/private";
import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';


export const actions = {
    default: async (event) => {
        try {
            const data = await event.request.formData();
            const valor = data.get("valor") as string; // Get the value from form data
            const contrasena = data.get("contrasena") as string; // Get the password from form data


            // Make a login request to the backend API
            const response = await fetch(`${API_URI}/login`, {
                method: 'POST',
                body: JSON.stringify({ valor, contrasena }),
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            console.log(response);

            // Handle response status
            if (!response.ok) {
                const errorData = await response.json();
                return fail(response.status, { error: errorData.error || 'Error desconocido' });
            }

            const login = await response.json();
            // Set the token in a cookie
            event.cookies.set('AuthorizationToken', `Bearer ${login.token}`, {
                httpOnly: true,
                path: '/',
                secure: true,
                sameSite: 'strict',
                maxAge: 60 * 60 * 2 // 2 hours
            });

            console.log('Bearer ' + login.token);

            return { success: "Auth successful" }; // Return success message

        } catch (error) {
            console.error('Error during login:', error);
            return fail(500, { error: 'Error processing request' }); // Return server error
        }
    },
} satisfies Actions;
