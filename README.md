# tower-of-hanoi
A go implementation of the Towers of Hanoi


This program solves the Towers of Hanoi using an iterative solution

## Execution
`go run tower.go -size=6`

## Example
```
$ go run tower.go -size=4
Size of the tower: 4
Left: 1:2:3:4:
Center: 
Right: 
Left: 2:3:4:
Center: 1:
Right: 
Left: 3:4:
Center: 1:
Right: 2:
Left: 3:4:
Center: 
Right: 1:2:
End of loop

Left: 4:
Center: 3:
Right: 1:2:
Left: 1:4:
Center: 3:
Right: 2:
Left: 1:4:
Center: 2:3:
Right: 
End of loop

Left: 4:
Center: 1:2:3:
Right: 
Left: 
Center: 1:2:3:
Right: 4:
Left: 
Center: 2:3:
Right: 1:4:
End of loop

Left: 2:
Center: 3:
Right: 1:4:
Left: 1:2:
Center: 3:
Right: 4:
Left: 1:2:
Center: 
Right: 3:4:
End of loop

Left: 2:
Center: 1:
Right: 3:4:
Left: 
Center: 1:
Right: 2:3:4:
Left: 
Center: 
Right: 1:2:3:4:
End of loop

Left: 
Center: 
Right: 1:2:3:4:
```

