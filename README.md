# Go Challenge: Temperature Converter

This project is a Go program that performs temperature conversions between Celsius, Fahrenheit, and Kelvin scales. The user can choose the input scale and get the converted values in the other two scales.

## How to run

1. Make sure you have Go installed on your system.
2. In the terminal, navigate to the project directory.
3. Run the command:

```bash
go run ConversionChallenge.go
```

## How to use

When you start the program, follow the instructions:
- Choose the input temperature scale:
  - 1 for Celsius
  - 2 for Fahrenheit
  - 3 for Kelvin
  - 4 to exit
- Enter the desired temperature value.
- The program will show the conversions to the other scales.
- Press Enter to continue or choose to exit.

## Example of use

```
Enter the scale do you want to use to convert the others (1-Celsius, 2-Fahrenheit, 3-Kelvin, 4- Leave):
1
Enter the temperature of water:
100
The temperature of water is:ºC100
This temperature in Fahrenheit is:ºF212
This temperature in Kelvin is :ºK373.15
Press Enter to continue...
```

## Code structure

- Functions for conversion between temperature scales:
  - Celsius to Fahrenheit/Kelvin
  - Fahrenheit to Celsius/Kelvin
  - Kelvin to Celsius/Fahrenheit
- Utility functions to clear the screen and pause execution.
- Main loop for user interaction.

## Requirements
- Go 1.13 or higher
- Compatible operating system (Windows, Linux, macOS)

## License
This project is for educational purposes only.
