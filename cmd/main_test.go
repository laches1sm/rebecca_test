package main

import (
    "testing"
)

func TestSayHello(t *testing.T) {
    // Arrange
    input := "Hello, Test!"
    expected := "Hello, Test!"

    // Act
    result := sayHello(input)

    // Assert
    if result != expected {
        t.Errorf("sayHello(%q) = %q; want %q", input, result, expected)
    }
}