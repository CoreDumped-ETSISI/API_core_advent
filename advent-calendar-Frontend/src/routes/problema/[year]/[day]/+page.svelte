<script lang="ts">
    import type { PageData, ActionData } from "../$types";
    export let data: PageData;

    import { enhance } from "$app/forms";

    export let form: ActionData; // This will hold the form response data
    const { problema, year, day } = data;

    let solution = '';

</script>

<main>
    <h1>Problema de Advent of Code - Año {year}, Día {day}</h1>

    {#if problema}
        <div>{@html problema}</div>
    {:else}
        <p>No hay problema disponible.</p>
    {/if}

    <form method="POST" use:enhance>
        <label for="solution">Your Solution:</label>
        <input type="text" id="solution" name="solution" bind:value={solution} required />
        <button type="submit">Submit</button>
    </form>

    {#if form?.error}
        <div class="error">
            ! {form.error}
        </div>
    {/if}

    {#if form?.result}
        <p>{form.result.message}</p>
    {/if}
</main>

<style>
    main {
        text-align: center;
        padding: 20px;
    }
    h1 {
        font-size: 2rem;
        color: #3498db;
    }
    div {
        margin-top: 20px;
        color: #333;
    }
    form {
        margin-top: 20px;
    }
</style>
