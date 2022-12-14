package day1

import java.nio.file.Files
import kotlin.io.path.Path

fun main() {
    val lines = Files.readAllLines(Path("day1/input.txt"))
    puzzle1(lines)
    puzzle2(lines)
}

private fun puzzle1(lines: List<String>) {
    var maxCalories = 0
    var current = 0
    lines.forEach {
        if (it.isBlank()) {
            if (current > maxCalories) {
                maxCalories = current
            }
            current = 0
        } else {
            current += it.toInt()
        }
    }
    println("puzzle1: $maxCalories")
}

private fun puzzle2(lines: List<String>) {
    val maxCaloriesList = mutableListOf(0, 0, 0)
    var current = 0
    lines.forEach {
        if (it.isBlank()) {
            maxCaloriesList.sort()
            if (current > maxCaloriesList.first()) {
                maxCaloriesList[0] = current
            }
            current = 0
        } else {
            current += it.toInt()
        }
    }
    println("puzzle2: ${maxCaloriesList.sum()}")
}
