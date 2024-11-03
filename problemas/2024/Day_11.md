Tetris en el saco de regalos


Ya queda menos para el dia-N (Navidad), toca organizar los regalos en el saco de CoreClaus. (Pensarás que va a ser el problema de la mochila pero no estarías más equivocado). El saco es mágico y para contener tantos regalos funciona con un conversor dimensional que cambia todos los regalos a paquetes cuadrados de 2 dimensiones. Los regalos se añaden a columnas concretas pero una vez caen se deslizan siguiendo las siguientes reglas:
- Por defecto un paquete se desliza a la izquierda.
- Si no puede moverse a la izquierda (porque hubiera ya un paquete allí) se desliza a la derecha
- Si se mueve encima de un agujero, cae al agujero.
- Una vez se ha deslizado o no ha podido deslizarse en absoluto, se fija en esa posición.

```
Teniendo el saco
5334
siendo cada número la altura de su columna

sería
☐
☐  ☐
☐☐☐☐
☐☐☐☐
☐☐☐☐

Si cae un paquete en la primera columna -> no se puede deslizar a la izquierda -> va a la derecha -> quedaría 5434
☐
☐☐ ☐
☐☐☐☐
☐☐☐☐
☐☐☐☐

Si cae uno en la tercera columna -> no se puede deslizar -> quedaría 5444
☐
☐☐☐☐
☐☐☐☐
☐☐☐☐
☐☐☐☐

Si cae uno en la tercera columna -> se desliza a la izquierda -> quedaría 5544
☐☐
☐☐☐☐
☐☐☐☐
☐☐☐☐
☐☐☐☐

```

El problema está en poder sacar los regalos del saco, para ello queremos saber las alturas de cada columna de regalos.
La solución es la cadena de todas las alturas (coge solo el último dígito => en 76 coge 6) del saco (en el ejemplo la respuesta sería 5544).