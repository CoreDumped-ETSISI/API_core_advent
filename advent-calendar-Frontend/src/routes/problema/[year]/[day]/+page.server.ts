// src/routes/problems/[year]/+page.server.ts
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad, Actions } from './$types';
import { API_URI } from "$env/static/private";

export const load: PageServerLoad = async (event) => {
    const auth = event.cookies.get("AuthorizationToken"); // Check for AuthorizationToken in cookies
    if (!auth) {
        throw redirect(303, "/login"); // Redirect to the admin page if token is not present
    }

    const { year, day } = event.params; // Get the year and day from the URL parameters

    try {
        const response = await fetch(`${API_URI}/${year}/${day}/`, {
            headers: {
                'Authorization': auth // Include the token in the request headers
            }
        });

        if (!response.ok) {
            throw new Error(`Failed to fetch: ${response.status}`);
        }

        const problema = await response.json(); // Fetch the problem data

        return { problema, year, day }; // Pass the data to the page
    } catch (error) {
        console.error("Error fetching problema:", error);
        return { problema: {}, year, day }; // Handle error by returning an empty object
    }
};



// Actions for handling POST requests
export const actions: Actions = {
    default: async (event) => {
        const auth = event.cookies.get("AuthorizationToken"); // Check for AuthorizationToken in cookies
        if (!auth) {
            throw redirect(303, "/login"); // Redirect to the login page if token is not present
        }

        const formData = await event.request.formData(); // Get the form data from the request

        const problemData = formData.get("solution"); // Replace with actual form field names


        try {
            const response = await fetch(`${API_URI}/${event.params.year}/${event.params.day}/`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': auth // Include the token in the request headers
                },
                body: JSON.stringify({ 
                    solucion_propuesta: problemData
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                return { error: errorData.error || 'Failed to submit the problem.' };
            }

            const result = await response.json();

            return { success: 'Problem submitted successfully!', result }; // Return success message and result
        } catch (error) {
            console.error("Error submitting problem:", error);
            return { error: 'Error processing your request.' }; // Handle errors
        }
    }
};