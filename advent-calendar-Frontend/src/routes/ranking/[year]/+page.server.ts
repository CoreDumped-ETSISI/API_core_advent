// src/routes/ranking/+page.server.ts

import type { PageServerLoad } from './$types';
import { API_URI } from "$env/static/private";


export const load: PageServerLoad = async (event ) => {
    const {year} = event.params; // Get the year from the URL params

    const response = await fetch(`${API_URI}/ranking/${year}`); // Adjust the URL to your API endpoint

    if (!response.ok) {
        return {
            status: response.status,
            body: { error: 'Error fetching ranking data' }
        };
    }

    const ranking = await response.json();

    return {
        ranking, year
    };
};
