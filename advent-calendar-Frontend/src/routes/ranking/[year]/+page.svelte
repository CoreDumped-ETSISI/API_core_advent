<!-- src/routes/ranking/+page.svelte -->
<script lang="ts">
    import type { PageData } from "../$types";

    export let data: PageData;
    const { ranking, year } = data;
    
    console.log(ranking, 'ranking');   
</script>

<h1>Ranking</h1>
<table>
    <thead>
        <tr>
            <th>Usuario</th>
            <th>Correctas</th>
            <th>Total Time Difference</th>
        </tr>
    </thead>
    <tbody>
        {#each ranking as user}
            <tr>
                <td>{user.usuario}</td>
                <td>{user.correctas}</td>
                <td>{user.total_time_difference}</td>
            </tr>
        {/each}
    </tbody>
</table>
