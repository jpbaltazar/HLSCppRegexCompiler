// Code generated from /home/zed4805/GolandProjects/thesisGoRemake/Resources/Grammar/PCRE.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 130, 628,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 4, 106,
	9, 106, 4, 107, 9, 107, 4, 108, 9, 108, 4, 109, 9, 109, 4, 110, 9, 110,
	4, 111, 9, 111, 4, 112, 9, 112, 4, 113, 9, 113, 4, 114, 9, 114, 4, 115,
	9, 115, 4, 116, 9, 116, 4, 117, 9, 117, 4, 118, 9, 118, 4, 119, 9, 119,
	4, 120, 9, 120, 4, 121, 9, 121, 4, 122, 9, 122, 4, 123, 9, 123, 4, 124,
	9, 124, 4, 125, 9, 125, 4, 126, 9, 126, 4, 127, 9, 127, 4, 128, 9, 128,
	4, 129, 9, 129, 4, 130, 9, 130, 4, 131, 9, 131, 4, 132, 9, 132, 4, 133,
	9, 133, 4, 134, 9, 134, 4, 135, 9, 135, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 285, 10, 6, 12, 6, 14,
	6, 288, 11, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 6, 15, 326, 10, 15, 13, 15, 14,
	15, 327, 3, 15, 3, 15, 5, 15, 332, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49,
	3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3,
	52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56,
	3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3,
	62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67,
	3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3,
	72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77,
	3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3,
	83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87, 3, 88,
	3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 3, 93, 3,
	93, 3, 94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97, 3, 98, 3, 98,
	3, 99, 3, 99, 3, 100, 3, 100, 3, 101, 3, 101, 3, 102, 3, 102, 3, 103, 3,
	103, 3, 104, 3, 104, 3, 105, 3, 105, 3, 106, 3, 106, 3, 107, 3, 107, 3,
	108, 3, 108, 3, 109, 3, 109, 3, 110, 3, 110, 3, 111, 3, 111, 3, 112, 3,
	112, 3, 113, 3, 113, 3, 114, 3, 114, 3, 115, 3, 115, 3, 116, 3, 116, 3,
	117, 3, 117, 3, 118, 3, 118, 3, 119, 3, 119, 3, 120, 3, 120, 3, 121, 3,
	121, 3, 122, 3, 122, 3, 123, 3, 123, 3, 124, 3, 124, 3, 125, 3, 125, 3,
	126, 3, 126, 3, 127, 3, 127, 3, 128, 3, 128, 3, 129, 3, 129, 3, 130, 3,
	130, 6, 130, 612, 10, 130, 13, 130, 14, 130, 613, 3, 131, 6, 131, 617,
	10, 131, 13, 131, 14, 131, 618, 3, 132, 3, 132, 3, 133, 3, 133, 3, 134,
	3, 134, 3, 135, 3, 135, 3, 286, 2, 136, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67,
	35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85,
	44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103,
	53, 105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119,
	61, 121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135,
	69, 137, 70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151,
	77, 153, 78, 155, 79, 157, 80, 159, 81, 161, 82, 163, 83, 165, 84, 167,
	85, 169, 86, 171, 87, 173, 88, 175, 89, 177, 90, 179, 91, 181, 92, 183,
	93, 185, 94, 187, 95, 189, 96, 191, 97, 193, 98, 195, 99, 197, 100, 199,
	101, 201, 102, 203, 103, 205, 104, 207, 105, 209, 106, 211, 107, 213, 108,
	215, 109, 217, 110, 219, 111, 221, 112, 223, 113, 225, 114, 227, 115, 229,
	116, 231, 117, 233, 118, 235, 119, 237, 120, 239, 121, 241, 122, 243, 123,
	245, 124, 247, 125, 249, 126, 251, 127, 253, 128, 255, 129, 257, 130, 259,
	2, 261, 2, 263, 2, 265, 2, 267, 2, 269, 2, 3, 2, 5, 5, 2, 50, 59, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 72, 99, 104, 3, 2, 2, 129, 2, 627, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3,
	2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65,
	3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2,
	73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2,
	2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2,
	2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2,
	2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103,
	3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2,
	2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3,
	2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2,
	125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2,
	2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139,
	3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2,
	2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3,
	2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2,
	161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3, 2,
	2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175,
	3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2,
	2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2, 2, 187, 3, 2, 2, 2, 2, 189, 3,
	2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193, 3, 2, 2, 2, 2, 195, 3, 2, 2, 2, 2,
	197, 3, 2, 2, 2, 2, 199, 3, 2, 2, 2, 2, 201, 3, 2, 2, 2, 2, 203, 3, 2,
	2, 2, 2, 205, 3, 2, 2, 2, 2, 207, 3, 2, 2, 2, 2, 209, 3, 2, 2, 2, 2, 211,
	3, 2, 2, 2, 2, 213, 3, 2, 2, 2, 2, 215, 3, 2, 2, 2, 2, 217, 3, 2, 2, 2,
	2, 219, 3, 2, 2, 2, 2, 221, 3, 2, 2, 2, 2, 223, 3, 2, 2, 2, 2, 225, 3,
	2, 2, 2, 2, 227, 3, 2, 2, 2, 2, 229, 3, 2, 2, 2, 2, 231, 3, 2, 2, 2, 2,
	233, 3, 2, 2, 2, 2, 235, 3, 2, 2, 2, 2, 237, 3, 2, 2, 2, 2, 239, 3, 2,
	2, 2, 2, 241, 3, 2, 2, 2, 2, 243, 3, 2, 2, 2, 2, 245, 3, 2, 2, 2, 2, 247,
	3, 2, 2, 2, 2, 249, 3, 2, 2, 2, 2, 251, 3, 2, 2, 2, 2, 253, 3, 2, 2, 2,
	2, 255, 3, 2, 2, 2, 2, 257, 3, 2, 2, 2, 3, 271, 3, 2, 2, 2, 5, 273, 3,
	2, 2, 2, 7, 275, 3, 2, 2, 2, 9, 277, 3, 2, 2, 2, 11, 280, 3, 2, 2, 2, 13,
	292, 3, 2, 2, 2, 15, 295, 3, 2, 2, 2, 17, 298, 3, 2, 2, 2, 19, 301, 3,
	2, 2, 2, 21, 304, 3, 2, 2, 2, 23, 307, 3, 2, 2, 2, 25, 310, 3, 2, 2, 2,
	27, 313, 3, 2, 2, 2, 29, 315, 3, 2, 2, 2, 31, 333, 3, 2, 2, 2, 33, 335,
	3, 2, 2, 2, 35, 338, 3, 2, 2, 2, 37, 341, 3, 2, 2, 2, 39, 344, 3, 2, 2,
	2, 41, 347, 3, 2, 2, 2, 43, 350, 3, 2, 2, 2, 45, 353, 3, 2, 2, 2, 47, 360,
	3, 2, 2, 2, 49, 367, 3, 2, 2, 2, 51, 370, 3, 2, 2, 2, 53, 373, 3, 2, 2,
	2, 55, 376, 3, 2, 2, 2, 57, 379, 3, 2, 2, 2, 59, 382, 3, 2, 2, 2, 61, 385,
	3, 2, 2, 2, 63, 388, 3, 2, 2, 2, 65, 391, 3, 2, 2, 2, 67, 393, 3, 2, 2,
	2, 69, 395, 3, 2, 2, 2, 71, 397, 3, 2, 2, 2, 73, 399, 3, 2, 2, 2, 75, 408,
	3, 2, 2, 2, 77, 418, 3, 2, 2, 2, 79, 420, 3, 2, 2, 2, 81, 422, 3, 2, 2,
	2, 83, 424, 3, 2, 2, 2, 85, 426, 3, 2, 2, 2, 87, 428, 3, 2, 2, 2, 89, 430,
	3, 2, 2, 2, 91, 433, 3, 2, 2, 2, 93, 436, 3, 2, 2, 2, 95, 439, 3, 2, 2,
	2, 97, 441, 3, 2, 2, 2, 99, 444, 3, 2, 2, 2, 101, 447, 3, 2, 2, 2, 103,
	450, 3, 2, 2, 2, 105, 453, 3, 2, 2, 2, 107, 456, 3, 2, 2, 2, 109, 459,
	3, 2, 2, 2, 111, 461, 3, 2, 2, 2, 113, 463, 3, 2, 2, 2, 115, 465, 3, 2,
	2, 2, 117, 467, 3, 2, 2, 2, 119, 469, 3, 2, 2, 2, 121, 471, 3, 2, 2, 2,
	123, 473, 3, 2, 2, 2, 125, 475, 3, 2, 2, 2, 127, 477, 3, 2, 2, 2, 129,
	479, 3, 2, 2, 2, 131, 481, 3, 2, 2, 2, 133, 483, 3, 2, 2, 2, 135, 485,
	3, 2, 2, 2, 137, 487, 3, 2, 2, 2, 139, 489, 3, 2, 2, 2, 141, 491, 3, 2,
	2, 2, 143, 493, 3, 2, 2, 2, 145, 495, 3, 2, 2, 2, 147, 497, 3, 2, 2, 2,
	149, 499, 3, 2, 2, 2, 151, 501, 3, 2, 2, 2, 153, 503, 3, 2, 2, 2, 155,
	505, 3, 2, 2, 2, 157, 507, 3, 2, 2, 2, 159, 509, 3, 2, 2, 2, 161, 511,
	3, 2, 2, 2, 163, 513, 3, 2, 2, 2, 165, 515, 3, 2, 2, 2, 167, 517, 3, 2,
	2, 2, 169, 519, 3, 2, 2, 2, 171, 521, 3, 2, 2, 2, 173, 523, 3, 2, 2, 2,
	175, 525, 3, 2, 2, 2, 177, 527, 3, 2, 2, 2, 179, 529, 3, 2, 2, 2, 181,
	531, 3, 2, 2, 2, 183, 533, 3, 2, 2, 2, 185, 535, 3, 2, 2, 2, 187, 537,
	3, 2, 2, 2, 189, 539, 3, 2, 2, 2, 191, 541, 3, 2, 2, 2, 193, 543, 3, 2,
	2, 2, 195, 545, 3, 2, 2, 2, 197, 547, 3, 2, 2, 2, 199, 549, 3, 2, 2, 2,
	201, 551, 3, 2, 2, 2, 203, 553, 3, 2, 2, 2, 205, 555, 3, 2, 2, 2, 207,
	557, 3, 2, 2, 2, 209, 559, 3, 2, 2, 2, 211, 561, 3, 2, 2, 2, 213, 563,
	3, 2, 2, 2, 215, 565, 3, 2, 2, 2, 217, 567, 3, 2, 2, 2, 219, 569, 3, 2,
	2, 2, 221, 571, 3, 2, 2, 2, 223, 573, 3, 2, 2, 2, 225, 575, 3, 2, 2, 2,
	227, 577, 3, 2, 2, 2, 229, 579, 3, 2, 2, 2, 231, 581, 3, 2, 2, 2, 233,
	583, 3, 2, 2, 2, 235, 585, 3, 2, 2, 2, 237, 587, 3, 2, 2, 2, 239, 589,
	3, 2, 2, 2, 241, 591, 3, 2, 2, 2, 243, 593, 3, 2, 2, 2, 245, 595, 3, 2,
	2, 2, 247, 597, 3, 2, 2, 2, 249, 599, 3, 2, 2, 2, 251, 601, 3, 2, 2, 2,
	253, 603, 3, 2, 2, 2, 255, 605, 3, 2, 2, 2, 257, 607, 3, 2, 2, 2, 259,
	611, 3, 2, 2, 2, 261, 616, 3, 2, 2, 2, 263, 620, 3, 2, 2, 2, 265, 622,
	3, 2, 2, 2, 267, 624, 3, 2, 2, 2, 269, 626, 3, 2, 2, 2, 271, 272, 7, 49,
	2, 2, 272, 4, 3, 2, 2, 2, 273, 274, 7, 15, 2, 2, 274, 6, 3, 2, 2, 2, 275,
	276, 7, 12, 2, 2, 276, 8, 3, 2, 2, 2, 277, 278, 7, 94, 2, 2, 278, 279,
	5, 265, 133, 2, 279, 10, 3, 2, 2, 2, 280, 281, 7, 94, 2, 2, 281, 282, 7,
	83, 2, 2, 282, 286, 3, 2, 2, 2, 283, 285, 11, 2, 2, 2, 284, 283, 3, 2,
	2, 2, 285, 288, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2,
	287, 289, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289, 290, 7, 94, 2, 2, 290,
	291, 7, 71, 2, 2, 291, 12, 3, 2, 2, 2, 292, 293, 7, 94, 2, 2, 293, 294,
	7, 99, 2, 2, 294, 14, 3, 2, 2, 2, 295, 296, 7, 94, 2, 2, 296, 297, 7, 101,
	2, 2, 297, 16, 3, 2, 2, 2, 298, 299, 7, 94, 2, 2, 299, 300, 7, 103, 2,
	2, 300, 18, 3, 2, 2, 2, 301, 302, 7, 94, 2, 2, 302, 303, 7, 104, 2, 2,
	303, 20, 3, 2, 2, 2, 304, 305, 7, 94, 2, 2, 305, 306, 7, 112, 2, 2, 306,
	22, 3, 2, 2, 2, 307, 308, 7, 94, 2, 2, 308, 309, 7, 116, 2, 2, 309, 24,
	3, 2, 2, 2, 310, 311, 7, 94, 2, 2, 311, 312, 7, 118, 2, 2, 312, 26, 3,
	2, 2, 2, 313, 314, 7, 94, 2, 2, 314, 28, 3, 2, 2, 2, 315, 316, 7, 94, 2,
	2, 316, 317, 7, 122, 2, 2, 317, 331, 3, 2, 2, 2, 318, 319, 5, 267, 134,
	2, 319, 320, 5, 267, 134, 2, 320, 332, 3, 2, 2, 2, 321, 322, 7, 125, 2,
	2, 322, 323, 5, 267, 134, 2, 323, 325, 5, 267, 134, 2, 324, 326, 5, 267,
	134, 2, 325, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 325, 3, 2, 2,
	2, 327, 328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 330, 7, 127, 2, 2,
	330, 332, 3, 2, 2, 2, 331, 318, 3, 2, 2, 2, 331, 321, 3, 2, 2, 2, 332,
	30, 3, 2, 2, 2, 333, 334, 7, 48, 2, 2, 334, 32, 3, 2, 2, 2, 335, 336, 7,
	94, 2, 2, 336, 337, 7, 69, 2, 2, 337, 34, 3, 2, 2, 2, 338, 339, 7, 94,
	2, 2, 339, 340, 7, 102, 2, 2, 340, 36, 3, 2, 2, 2, 341, 342, 7, 94, 2,
	2, 342, 343, 7, 70, 2, 2, 343, 38, 3, 2, 2, 2, 344, 345, 7, 94, 2, 2, 345,
	346, 7, 106, 2, 2, 346, 40, 3, 2, 2, 2, 347, 348, 7, 94, 2, 2, 348, 349,
	7, 74, 2, 2, 349, 42, 3, 2, 2, 2, 350, 351, 7, 94, 2, 2, 351, 352, 7, 80,
	2, 2, 352, 44, 3, 2, 2, 2, 353, 354, 7, 94, 2, 2, 354, 355, 7, 114, 2,
	2, 355, 356, 7, 125, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 5, 259, 130,
	2, 358, 359, 7, 127, 2, 2, 359, 46, 3, 2, 2, 2, 360, 361, 7, 94, 2, 2,
	361, 362, 7, 82, 2, 2, 362, 363, 7, 125, 2, 2, 363, 364, 3, 2, 2, 2, 364,
	365, 5, 259, 130, 2, 365, 366, 7, 127, 2, 2, 366, 48, 3, 2, 2, 2, 367,
	368, 7, 94, 2, 2, 368, 369, 7, 84, 2, 2, 369, 50, 3, 2, 2, 2, 370, 371,
	7, 94, 2, 2, 371, 372, 7, 117, 2, 2, 372, 52, 3, 2, 2, 2, 373, 374, 7,
	94, 2, 2, 374, 375, 7, 85, 2, 2, 375, 54, 3, 2, 2, 2, 376, 377, 7, 94,
	2, 2, 377, 378, 7, 120, 2, 2, 378, 56, 3, 2, 2, 2, 379, 380, 7, 94, 2,
	2, 380, 381, 7, 88, 2, 2, 381, 58, 3, 2, 2, 2, 382, 383, 7, 94, 2, 2, 383,
	384, 7, 121, 2, 2, 384, 60, 3, 2, 2, 2, 385, 386, 7, 94, 2, 2, 386, 387,
	7, 89, 2, 2, 387, 62, 3, 2, 2, 2, 388, 389, 7, 94, 2, 2, 389, 390, 7, 90,
	2, 2, 390, 64, 3, 2, 2, 2, 391, 392, 7, 93, 2, 2, 392, 66, 3, 2, 2, 2,
	393, 394, 7, 95, 2, 2, 394, 68, 3, 2, 2, 2, 395, 396, 7, 96, 2, 2, 396,
	70, 3, 2, 2, 2, 397, 398, 7, 47, 2, 2, 398, 72, 3, 2, 2, 2, 399, 400, 7,
	93, 2, 2, 400, 401, 7, 93, 2, 2, 401, 402, 7, 60, 2, 2, 402, 403, 3, 2,
	2, 2, 403, 404, 5, 261, 131, 2, 404, 405, 7, 60, 2, 2, 405, 406, 7, 95,
	2, 2, 406, 407, 7, 95, 2, 2, 407, 74, 3, 2, 2, 2, 408, 409, 7, 93, 2, 2,
	409, 410, 7, 93, 2, 2, 410, 411, 7, 60, 2, 2, 411, 412, 7, 96, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 414, 5, 261, 131, 2, 414, 415, 7, 60, 2, 2, 415,
	416, 7, 95, 2, 2, 416, 417, 7, 95, 2, 2, 417, 76, 3, 2, 2, 2, 418, 419,
	7, 65, 2, 2, 419, 78, 3, 2, 2, 2, 420, 421, 7, 45, 2, 2, 421, 80, 3, 2,
	2, 2, 422, 423, 7, 44, 2, 2, 423, 82, 3, 2, 2, 2, 424, 425, 7, 125, 2,
	2, 425, 84, 3, 2, 2, 2, 426, 427, 7, 127, 2, 2, 427, 86, 3, 2, 2, 2, 428,
	429, 7, 46, 2, 2, 429, 88, 3, 2, 2, 2, 430, 431, 7, 94, 2, 2, 431, 432,
	7, 100, 2, 2, 432, 90, 3, 2, 2, 2, 433, 434, 7, 94, 2, 2, 434, 435, 7,
	68, 2, 2, 435, 92, 3, 2, 2, 2, 436, 437, 7, 94, 2, 2, 437, 438, 7, 67,
	2, 2, 438, 94, 3, 2, 2, 2, 439, 440, 7, 38, 2, 2, 440, 96, 3, 2, 2, 2,
	441, 442, 7, 94, 2, 2, 442, 443, 7, 92, 2, 2, 443, 98, 3, 2, 2, 2, 444,
	445, 7, 94, 2, 2, 445, 446, 7, 124, 2, 2, 446, 100, 3, 2, 2, 2, 447, 448,
	7, 94, 2, 2, 448, 449, 7, 73, 2, 2, 449, 102, 3, 2, 2, 2, 450, 451, 7,
	94, 2, 2, 451, 452, 7, 77, 2, 2, 452, 104, 3, 2, 2, 2, 453, 454, 7, 94,
	2, 2, 454, 455, 7, 105, 2, 2, 455, 106, 3, 2, 2, 2, 456, 457, 7, 94, 2,
	2, 457, 458, 7, 109, 2, 2, 458, 108, 3, 2, 2, 2, 459, 460, 7, 126, 2, 2,
	460, 110, 3, 2, 2, 2, 461, 462, 7, 42, 2, 2, 462, 112, 3, 2, 2, 2, 463,
	464, 7, 43, 2, 2, 464, 114, 3, 2, 2, 2, 465, 466, 7, 62, 2, 2, 466, 116,
	3, 2, 2, 2, 467, 468, 7, 64, 2, 2, 468, 118, 3, 2, 2, 2, 469, 470, 7, 41,
	2, 2, 470, 120, 3, 2, 2, 2, 471, 472, 7, 97, 2, 2, 472, 122, 3, 2, 2, 2,
	473, 474, 7, 60, 2, 2, 474, 124, 3, 2, 2, 2, 475, 476, 7, 37, 2, 2, 476,
	126, 3, 2, 2, 2, 477, 478, 7, 63, 2, 2, 478, 128, 3, 2, 2, 2, 479, 480,
	7, 35, 2, 2, 480, 130, 3, 2, 2, 2, 481, 482, 7, 40, 2, 2, 482, 132, 3,
	2, 2, 2, 483, 484, 7, 99, 2, 2, 484, 134, 3, 2, 2, 2, 485, 486, 7, 100,
	2, 2, 486, 136, 3, 2, 2, 2, 487, 488, 7, 101, 2, 2, 488, 138, 3, 2, 2,
	2, 489, 490, 7, 102, 2, 2, 490, 140, 3, 2, 2, 2, 491, 492, 7, 103, 2, 2,
	492, 142, 3, 2, 2, 2, 493, 494, 7, 104, 2, 2, 494, 144, 3, 2, 2, 2, 495,
	496, 7, 105, 2, 2, 496, 146, 3, 2, 2, 2, 497, 498, 7, 106, 2, 2, 498, 148,
	3, 2, 2, 2, 499, 500, 7, 107, 2, 2, 500, 150, 3, 2, 2, 2, 501, 502, 7,
	108, 2, 2, 502, 152, 3, 2, 2, 2, 503, 504, 7, 109, 2, 2, 504, 154, 3, 2,
	2, 2, 505, 506, 7, 110, 2, 2, 506, 156, 3, 2, 2, 2, 507, 508, 7, 111, 2,
	2, 508, 158, 3, 2, 2, 2, 509, 510, 7, 112, 2, 2, 510, 160, 3, 2, 2, 2,
	511, 512, 7, 113, 2, 2, 512, 162, 3, 2, 2, 2, 513, 514, 7, 114, 2, 2, 514,
	164, 3, 2, 2, 2, 515, 516, 7, 115, 2, 2, 516, 166, 3, 2, 2, 2, 517, 518,
	7, 116, 2, 2, 518, 168, 3, 2, 2, 2, 519, 520, 7, 117, 2, 2, 520, 170, 3,
	2, 2, 2, 521, 522, 7, 118, 2, 2, 522, 172, 3, 2, 2, 2, 523, 524, 7, 119,
	2, 2, 524, 174, 3, 2, 2, 2, 525, 526, 7, 120, 2, 2, 526, 176, 3, 2, 2,
	2, 527, 528, 7, 121, 2, 2, 528, 178, 3, 2, 2, 2, 529, 530, 7, 122, 2, 2,
	530, 180, 3, 2, 2, 2, 531, 532, 7, 123, 2, 2, 532, 182, 3, 2, 2, 2, 533,
	534, 7, 124, 2, 2, 534, 184, 3, 2, 2, 2, 535, 536, 7, 67, 2, 2, 536, 186,
	3, 2, 2, 2, 537, 538, 7, 68, 2, 2, 538, 188, 3, 2, 2, 2, 539, 540, 7, 69,
	2, 2, 540, 190, 3, 2, 2, 2, 541, 542, 7, 70, 2, 2, 542, 192, 3, 2, 2, 2,
	543, 544, 7, 71, 2, 2, 544, 194, 3, 2, 2, 2, 545, 546, 7, 72, 2, 2, 546,
	196, 3, 2, 2, 2, 547, 548, 7, 73, 2, 2, 548, 198, 3, 2, 2, 2, 549, 550,
	7, 74, 2, 2, 550, 200, 3, 2, 2, 2, 551, 552, 7, 75, 2, 2, 552, 202, 3,
	2, 2, 2, 553, 554, 7, 76, 2, 2, 554, 204, 3, 2, 2, 2, 555, 556, 7, 77,
	2, 2, 556, 206, 3, 2, 2, 2, 557, 558, 7, 78, 2, 2, 558, 208, 3, 2, 2, 2,
	559, 560, 7, 79, 2, 2, 560, 210, 3, 2, 2, 2, 561, 562, 7, 80, 2, 2, 562,
	212, 3, 2, 2, 2, 563, 564, 7, 81, 2, 2, 564, 214, 3, 2, 2, 2, 565, 566,
	7, 82, 2, 2, 566, 216, 3, 2, 2, 2, 567, 568, 7, 83, 2, 2, 568, 218, 3,
	2, 2, 2, 569, 570, 7, 84, 2, 2, 570, 220, 3, 2, 2, 2, 571, 572, 7, 85,
	2, 2, 572, 222, 3, 2, 2, 2, 573, 574, 7, 86, 2, 2, 574, 224, 3, 2, 2, 2,
	575, 576, 7, 87, 2, 2, 576, 226, 3, 2, 2, 2, 577, 578, 7, 88, 2, 2, 578,
	228, 3, 2, 2, 2, 579, 580, 7, 89, 2, 2, 580, 230, 3, 2, 2, 2, 581, 582,
	7, 90, 2, 2, 582, 232, 3, 2, 2, 2, 583, 584, 7, 91, 2, 2, 584, 234, 3,
	2, 2, 2, 585, 586, 7, 92, 2, 2, 586, 236, 3, 2, 2, 2, 587, 588, 7, 51,
	2, 2, 588, 238, 3, 2, 2, 2, 589, 590, 7, 52, 2, 2, 590, 240, 3, 2, 2, 2,
	591, 592, 7, 53, 2, 2, 592, 242, 3, 2, 2, 2, 593, 594, 7, 54, 2, 2, 594,
	244, 3, 2, 2, 2, 595, 596, 7, 55, 2, 2, 596, 246, 3, 2, 2, 2, 597, 598,
	7, 56, 2, 2, 598, 248, 3, 2, 2, 2, 599, 600, 7, 57, 2, 2, 600, 250, 3,
	2, 2, 2, 601, 602, 7, 58, 2, 2, 602, 252, 3, 2, 2, 2, 603, 604, 7, 59,
	2, 2, 604, 254, 3, 2, 2, 2, 605, 606, 7, 50, 2, 2, 606, 256, 3, 2, 2, 2,
	607, 608, 11, 2, 2, 2, 608, 258, 3, 2, 2, 2, 609, 612, 7, 97, 2, 2, 610,
	612, 5, 263, 132, 2, 611, 609, 3, 2, 2, 2, 611, 610, 3, 2, 2, 2, 612, 613,
	3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614, 260, 3, 2,
	2, 2, 615, 617, 5, 263, 132, 2, 616, 615, 3, 2, 2, 2, 617, 618, 3, 2, 2,
	2, 618, 616, 3, 2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 262, 3, 2, 2, 2, 620,
	621, 9, 2, 2, 2, 621, 264, 3, 2, 2, 2, 622, 623, 10, 2, 2, 2, 623, 266,
	3, 2, 2, 2, 624, 625, 9, 3, 2, 2, 625, 268, 3, 2, 2, 2, 626, 627, 9, 4,
	2, 2, 627, 270, 3, 2, 2, 2, 9, 2, 286, 327, 331, 611, 613, 618, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'/'", "'\r'", "'\n'", "", "", "'\\a'", "'\\c'", "'\\e'", "'\\f'",
	"'\\n'", "'\\r'", "'\\t'", "'\\'", "", "'.'", "'\\C'", "'\\d'", "'\\D'",
	"'\\h'", "'\\H'", "'\\N'", "", "", "'\\R'", "'\\s'", "'\\S'", "'\\v'",
	"'\\V'", "'\\w'", "'\\W'", "'\\X'", "'['", "']'", "'^'", "'-'", "", "",
	"'?'", "'+'", "'*'", "'{'", "'}'", "','", "'\\b'", "'\\B'", "'\\A'", "'$'",
	"'\\Z'", "'\\z'", "'\\G'", "'\\K'", "'\\g'", "'\\k'", "'|'", "'('", "')'",
	"'<'", "'>'", "'''", "'_'", "':'", "'#'", "'='", "'!'", "'&'", "'a'", "'b'",
	"'c'", "'d'", "'e'", "'f'", "'g'", "'h'", "'i'", "'j'", "'k'", "'l'", "'m'",
	"'n'", "'o'", "'p'", "'q'", "'r'", "'s'", "'t'", "'u'", "'v'", "'w'", "'x'",
	"'y'", "'z'", "'A'", "'B'", "'C'", "'D'", "'E'", "'F'", "'G'", "'H'", "'I'",
	"'J'", "'K'", "'L'", "'M'", "'N'", "'O'", "'P'", "'Q'", "'R'", "'S'", "'T'",
	"'U'", "'V'", "'W'", "'X'", "'Y'", "'Z'", "'1'", "'2'", "'3'", "'4'", "'5'",
	"'6'", "'7'", "'8'", "'9'", "'0'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "Quoted", "BlockQuoted", "BellChar", "ControlChar", "EscapeChar",
	"FormFeed", "NewLine", "CarriageReturn", "Tab", "Backslash", "HexChar",
	"Dot", "OneDataUnit", "DecimalDigit", "NotDecimalDigit", "HorizontalWhiteSpace",
	"NotHorizontalWhiteSpace", "NotNewLine", "CharWithProperty", "CharWithoutProperty",
	"NewLineSequence", "WhiteSpace", "NotWhiteSpace", "VerticalWhiteSpace",
	"NotVerticalWhiteSpace", "WordChar", "NotWordChar", "ExtendedUnicodeChar",
	"CharacterClassStart", "CharacterClassEnd", "Caret", "Hyphen", "POSIXNamedSet",
	"POSIXNegatedNamedSet", "QuestionMark", "Plus", "Star", "OpenBrace", "CloseBrace",
	"Comma", "WordBoundary", "NonWordBoundary", "StartOfSubject", "EndOfSubjectOrLine",
	"EndOfSubjectOrLineEndOfSubject", "EndOfSubject", "PreviousMatchInSubject",
	"ResetStartMatch", "SubroutineOrNamedReferenceStartG", "NamedReferenceStartK",
	"Pipe", "OpenParen", "CloseParen", "LessThan", "GreaterThan", "SingleQuote",
	"Underscore", "Colon", "Hash", "Equals", "Exclamation", "Ampersand", "ALC",
	"BLC", "CLC", "DLC", "ELC", "FLC", "GLC", "HLC", "ILC", "JLC", "KLC", "LLC",
	"MLC", "NLC", "OLC", "PLC", "QLC", "RLC", "SLC", "TLC", "ULC", "VLC", "WLC",
	"XLC", "YLC", "ZLC", "AUC", "BUC", "CUC", "DUC", "EUC", "FUC", "GUC", "HUC",
	"IUC", "JUC", "KUC", "LUC", "MUC", "NUC", "OUC", "PUC", "QUC", "RUC", "SUC",
	"TUC", "UUC", "VUC", "WUC", "XUC", "YUC", "ZUC", "D1", "D2", "D3", "D4",
	"D5", "D6", "D7", "D8", "D9", "D0", "OtherChar",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "Quoted", "BlockQuoted", "BellChar", "ControlChar",
	"EscapeChar", "FormFeed", "NewLine", "CarriageReturn", "Tab", "Backslash",
	"HexChar", "Dot", "OneDataUnit", "DecimalDigit", "NotDecimalDigit", "HorizontalWhiteSpace",
	"NotHorizontalWhiteSpace", "NotNewLine", "CharWithProperty", "CharWithoutProperty",
	"NewLineSequence", "WhiteSpace", "NotWhiteSpace", "VerticalWhiteSpace",
	"NotVerticalWhiteSpace", "WordChar", "NotWordChar", "ExtendedUnicodeChar",
	"CharacterClassStart", "CharacterClassEnd", "Caret", "Hyphen", "POSIXNamedSet",
	"POSIXNegatedNamedSet", "QuestionMark", "Plus", "Star", "OpenBrace", "CloseBrace",
	"Comma", "WordBoundary", "NonWordBoundary", "StartOfSubject", "EndOfSubjectOrLine",
	"EndOfSubjectOrLineEndOfSubject", "EndOfSubject", "PreviousMatchInSubject",
	"ResetStartMatch", "SubroutineOrNamedReferenceStartG", "NamedReferenceStartK",
	"Pipe", "OpenParen", "CloseParen", "LessThan", "GreaterThan", "SingleQuote",
	"Underscore", "Colon", "Hash", "Equals", "Exclamation", "Ampersand", "ALC",
	"BLC", "CLC", "DLC", "ELC", "FLC", "GLC", "HLC", "ILC", "JLC", "KLC", "LLC",
	"MLC", "NLC", "OLC", "PLC", "QLC", "RLC", "SLC", "TLC", "ULC", "VLC", "WLC",
	"XLC", "YLC", "ZLC", "AUC", "BUC", "CUC", "DUC", "EUC", "FUC", "GUC", "HUC",
	"IUC", "JUC", "KUC", "LUC", "MUC", "NUC", "OUC", "PUC", "QUC", "RUC", "SUC",
	"TUC", "UUC", "VUC", "WUC", "XUC", "YUC", "ZUC", "D1", "D2", "D3", "D4",
	"D5", "D6", "D7", "D8", "D9", "D0", "OtherChar", "UnderscoreAlphaNumerics",
	"AlphaNumerics", "AlphaNumeric", "NonAlphaNumeric", "HexDigit", "ASCII",
}

type PCRELexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPCRELexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PCRELexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPCRELexer(input antlr.CharStream) *PCRELexer {
	l := new(PCRELexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PCRE.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PCRELexer tokens.
const (
	PCRELexerT__0                             = 1
	PCRELexerT__1                             = 2
	PCRELexerT__2                             = 3
	PCRELexerQuoted                           = 4
	PCRELexerBlockQuoted                      = 5
	PCRELexerBellChar                         = 6
	PCRELexerControlChar                      = 7
	PCRELexerEscapeChar                       = 8
	PCRELexerFormFeed                         = 9
	PCRELexerNewLine                          = 10
	PCRELexerCarriageReturn                   = 11
	PCRELexerTab                              = 12
	PCRELexerBackslash                        = 13
	PCRELexerHexChar                          = 14
	PCRELexerDot                              = 15
	PCRELexerOneDataUnit                      = 16
	PCRELexerDecimalDigit                     = 17
	PCRELexerNotDecimalDigit                  = 18
	PCRELexerHorizontalWhiteSpace             = 19
	PCRELexerNotHorizontalWhiteSpace          = 20
	PCRELexerNotNewLine                       = 21
	PCRELexerCharWithProperty                 = 22
	PCRELexerCharWithoutProperty              = 23
	PCRELexerNewLineSequence                  = 24
	PCRELexerWhiteSpace                       = 25
	PCRELexerNotWhiteSpace                    = 26
	PCRELexerVerticalWhiteSpace               = 27
	PCRELexerNotVerticalWhiteSpace            = 28
	PCRELexerWordChar                         = 29
	PCRELexerNotWordChar                      = 30
	PCRELexerExtendedUnicodeChar              = 31
	PCRELexerCharacterClassStart              = 32
	PCRELexerCharacterClassEnd                = 33
	PCRELexerCaret                            = 34
	PCRELexerHyphen                           = 35
	PCRELexerPOSIXNamedSet                    = 36
	PCRELexerPOSIXNegatedNamedSet             = 37
	PCRELexerQuestionMark                     = 38
	PCRELexerPlus                             = 39
	PCRELexerStar                             = 40
	PCRELexerOpenBrace                        = 41
	PCRELexerCloseBrace                       = 42
	PCRELexerComma                            = 43
	PCRELexerWordBoundary                     = 44
	PCRELexerNonWordBoundary                  = 45
	PCRELexerStartOfSubject                   = 46
	PCRELexerEndOfSubjectOrLine               = 47
	PCRELexerEndOfSubjectOrLineEndOfSubject   = 48
	PCRELexerEndOfSubject                     = 49
	PCRELexerPreviousMatchInSubject           = 50
	PCRELexerResetStartMatch                  = 51
	PCRELexerSubroutineOrNamedReferenceStartG = 52
	PCRELexerNamedReferenceStartK             = 53
	PCRELexerPipe                             = 54
	PCRELexerOpenParen                        = 55
	PCRELexerCloseParen                       = 56
	PCRELexerLessThan                         = 57
	PCRELexerGreaterThan                      = 58
	PCRELexerSingleQuote                      = 59
	PCRELexerUnderscore                       = 60
	PCRELexerColon                            = 61
	PCRELexerHash                             = 62
	PCRELexerEquals                           = 63
	PCRELexerExclamation                      = 64
	PCRELexerAmpersand                        = 65
	PCRELexerALC                              = 66
	PCRELexerBLC                              = 67
	PCRELexerCLC                              = 68
	PCRELexerDLC                              = 69
	PCRELexerELC                              = 70
	PCRELexerFLC                              = 71
	PCRELexerGLC                              = 72
	PCRELexerHLC                              = 73
	PCRELexerILC                              = 74
	PCRELexerJLC                              = 75
	PCRELexerKLC                              = 76
	PCRELexerLLC                              = 77
	PCRELexerMLC                              = 78
	PCRELexerNLC                              = 79
	PCRELexerOLC                              = 80
	PCRELexerPLC                              = 81
	PCRELexerQLC                              = 82
	PCRELexerRLC                              = 83
	PCRELexerSLC                              = 84
	PCRELexerTLC                              = 85
	PCRELexerULC                              = 86
	PCRELexerVLC                              = 87
	PCRELexerWLC                              = 88
	PCRELexerXLC                              = 89
	PCRELexerYLC                              = 90
	PCRELexerZLC                              = 91
	PCRELexerAUC                              = 92
	PCRELexerBUC                              = 93
	PCRELexerCUC                              = 94
	PCRELexerDUC                              = 95
	PCRELexerEUC                              = 96
	PCRELexerFUC                              = 97
	PCRELexerGUC                              = 98
	PCRELexerHUC                              = 99
	PCRELexerIUC                              = 100
	PCRELexerJUC                              = 101
	PCRELexerKUC                              = 102
	PCRELexerLUC                              = 103
	PCRELexerMUC                              = 104
	PCRELexerNUC                              = 105
	PCRELexerOUC                              = 106
	PCRELexerPUC                              = 107
	PCRELexerQUC                              = 108
	PCRELexerRUC                              = 109
	PCRELexerSUC                              = 110
	PCRELexerTUC                              = 111
	PCRELexerUUC                              = 112
	PCRELexerVUC                              = 113
	PCRELexerWUC                              = 114
	PCRELexerXUC                              = 115
	PCRELexerYUC                              = 116
	PCRELexerZUC                              = 117
	PCRELexerD1                               = 118
	PCRELexerD2                               = 119
	PCRELexerD3                               = 120
	PCRELexerD4                               = 121
	PCRELexerD5                               = 122
	PCRELexerD6                               = 123
	PCRELexerD7                               = 124
	PCRELexerD8                               = 125
	PCRELexerD9                               = 126
	PCRELexerD0                               = 127
	PCRELexerOtherChar                        = 128
)