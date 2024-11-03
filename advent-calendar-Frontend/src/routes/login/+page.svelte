<script lang="ts">
    import { enhance } from "$app/forms";
    import { goto } from '$app/navigation';
    import { page } from '$app/stores';
    import type { ActionData } from "./$types";
    
    export let form: ActionData; // This will hold the form response data

    // Reactive statement to handle navigation after successful login
    $: if (form?.success) {
        goto('/'); // Redirect to the home page on success
    }
</script>


<div class="center">
    <div class="container">
        <img src="/logo.png" alt="logo" />
        <form method="POST" use:enhance>
            <input
                name="valor" 
                type="text"
                id="valor"
                placeholder="Correo o Usuario"
                required
            />
            <input
                name="contrasena" 
                type="password"
                id="contrasena"
                placeholder="Contraseña"
                required
            />
            <button type="submit">Iniciar Sesión</button>
            {#if form?.error}
                <div class="error">
                    ! {form.error}
                </div>
            {/if}
        </form>
    </div>
</div>

<style>
    .center {
        display: flex;
        justify-content: center;
        align-items: center;
    }
    img {
        max-width: 200px;
    }
    input {
        height: max-content;
        border: transparent;
        font-size: 16px;
        border-bottom: 1px solid #5fb030;
        background-color: transparent;
        padding: 4px;
    }
    input:focus {
        border-radius: 8px;
        border-bottom: 1px solid transparent;
        outline: 2px solid #5fb030;
    }

    .container {
        display: flex;
        flex-direction: column;
        gap: 8px;
        border-radius: 24px;
        padding: 24px;
        max-height: fit-content;
        background: white;
        justify-content: center;
        align-items: center;
        box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.3);
    }
    button {
        width: fit-content;
        background-color: #5fb030;
        color: white;
        box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.205);
        border-radius: 16px;
        border: transparent;
        padding: 6px;
    }
    button:active {
        box-shadow: none;
    }
    .error {
        color: red;
    }
</style>
