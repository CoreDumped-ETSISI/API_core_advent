import { API_URI } from "$env/static/private";
import type { RequestHandler } from '@sveltejs/kit';

export const load: RequestHandler = async ({ params, request }) => {
    const { year } = params;

    // Retrieve the token from the Authorization header

    try {
        const response = await fetch(`${API_URI}/${year}/`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
            
        });

        if (!response.ok) {
            throw new Error(`Failed to fetch: ${response.status}`);
        }

        const problemas = await response.json();
        return { problemas, year };

    } catch (error) {
        console.log(error);
        console.error("Error fetching problemas:", error);
        return { problemas: [], year, message: 'Error fetching data' };
    }
};
