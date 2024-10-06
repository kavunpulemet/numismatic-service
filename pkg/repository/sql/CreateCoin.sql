INSERT INTO coins (name, country, year, denomination, material, weight, diameter, thickness, condition, mintmark, historicalinfo, value)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id