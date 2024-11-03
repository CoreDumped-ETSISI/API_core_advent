// +page.server.ts
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ request }) => {
    try {
        const { valor, contrasena } = await request.json();

        // Make a login request to the backend API
        const response = await fetch('http://localhost:8000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ valor, contrasena }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            return new Response(JSON.stringify({ error: errorData.error || 'Error desconocido' }), { status: 401 });
        }

        const data = await response.json();
        return new Response(JSON.stringify({ token: data.token }), { status: 200 });
    } catch (error) {
        console.error('Error during login:', error);
        return new Response(JSON.stringify({ error: 'Error processing request' }), { status: 500 });
    }
};
