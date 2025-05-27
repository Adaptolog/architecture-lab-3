#!/bin/bash
# Початкова позиція
x=0.5
y=0.5

# Створюємо фігуру в початковій позиції
curl -X POST http://localhost:17000 -d "figure $x $y"
curl -X POST http://localhost:17000 -d "update"
sleep 1

# Переміщуємо фігуру крок за кроком (1% за крок)
for i in {1..20}; do
    curl -X POST http://localhost:17000 -d "move 0.01 0.01"
    curl -X POST http://localhost:17000 -d "update"
    sleep 1
done