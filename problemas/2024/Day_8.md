Core Dumped en el CoreTrineo
https://adventofcode.com/2020/day/8

CoreClaus está probando su trineo antes de Navidad. Sin embargo el sistema operativo del trineo ha fallado y ha devuelto un `segmentation fault (Core Dumped)`. Has localizado que el problema es un bucle infinito en el código de inicialización del trineo. El código esta escrito en un tipo de ensamblador muy antiguo y mágico. Has conseguido un manual polvoriento con la siguiente información:

- `acc` aumenta o disminuye una única variable global llamadoa acumulador por el valor dado en el argumento. Por ejemplo, `acc +7` aumentaría el acumulador en 7. El acumulador se inicializa en 0. Después de una instrucción `acc`, se ejecuta la instrucción inmediatamente inferior.
- `jmp` salta a una nueva instrucción relativa a sí misma. La siguiente instrucción a ejecutar se encuentra usando el argumento como un offset desde la instrucción `jmp`; por ejemplo `jmp +2` saltaría a la siguiente instrucción, `jmp +1` continuaría a la instrucción inmediatamente inferior, y `jmp -20` haría que la instrucción 20 líneas por encima fuera la siguiente en ejecutarse.
- `beq` funciona igual que `jump` pero solo salta si el acumulador es igual a 0.
- `nop` significa No Operation - no hace nada. La instrucción inmediatamente inferior se ejecuta a continuación.

No sabes exactamente donde ocurre el bucle, para averiguarla, quieres saber cual es el valor del acumulador justo antes de que se ejecute una instrucción por segunda vez.