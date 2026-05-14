1. Pitch Shift (El más importante)
Es la base del sonido "Screwed". Baja la tonalidad de la voz o del instrumento sin necesariamente cambiar la velocidad (aunque en el estilo clásico de Houston suelen ir de la mano).

Valor en JSON: {"pitch": -5} (semitonos).

2. Reverb (Espacialidad)
Fundamental para darle esa atmósfera "espacial" y profesional a los samples generados por IA, que a veces pueden sonar muy "secos".

Valor en JSON: {"reverb_dry_wet": 0.5, "reverb_room_size": 0.8}.

3. Bitcrush (Textura Lo-Fi)
Reduce la calidad del audio intencionalmente para que suene como un sampler antiguo de los 80s o 90s. Ideal para samples de baterías o melodías melancólicas.

Valor en JSON: {"bit_depth": 8, "sample_rate_reduction": 0.4}.

4. Delay (Eco)
Añade repeticiones rítmicas. Muy usado en transiciones.

Valor en JSON: {"delay_time": "1/4", "feedback": 0.3}.

5. Low Pass Filter (Filtro de Bajos)
Corta las frecuencias agudas para que el sample suene como si estuviera "bajo el agua" o escuchándose desde otra habitación.

Valor en JSON: {"cutoff_freq": 1000} (Hz).