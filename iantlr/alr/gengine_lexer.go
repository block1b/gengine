// Code generated from /Users/renyunyi/go/src/gengine/iantlr/gengine.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 45, 424, 
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
	70, 4, 71, 9, 71, 4, 72, 9, 72, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 
	27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 
	3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 5, 34, 223, 10, 34, 3, 34, 6, 34, 226, 
	10, 34, 13, 34, 14, 34, 227, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 
	3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 
	42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 
	3, 44, 3, 44, 3, 45, 6, 45, 281, 10, 45, 13, 45, 14, 45, 282, 3, 45, 7, 
	45, 286, 10, 45, 12, 45, 14, 45, 289, 11, 45, 3, 46, 6, 46, 292, 10, 46, 
	13, 46, 14, 46, 293, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 
	50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 
	3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 
	58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 
	3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 
	68, 3, 68, 3, 68, 3, 68, 7, 68, 349, 10, 68, 12, 68, 14, 68, 352, 11, 68, 
	3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 6, 70, 361, 10, 70, 13, 
	70, 14, 70, 362, 5, 70, 365, 10, 70, 3, 70, 3, 70, 6, 70, 369, 10, 70, 
	13, 70, 14, 70, 370, 3, 70, 6, 70, 374, 10, 70, 13, 70, 14, 70, 375, 3, 
	70, 3, 70, 3, 70, 3, 70, 6, 70, 382, 10, 70, 13, 70, 14, 70, 383, 5, 70, 
	386, 10, 70, 3, 70, 3, 70, 6, 70, 390, 10, 70, 13, 70, 14, 70, 391, 3, 
	70, 3, 70, 3, 70, 6, 70, 397, 10, 70, 13, 70, 14, 70, 398, 3, 70, 3, 70, 
	5, 70, 403, 10, 70, 3, 71, 3, 71, 3, 71, 3, 71, 7, 71, 409, 10, 71, 12, 
	71, 14, 71, 412, 11, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 72, 6, 72, 419, 
	10, 72, 13, 72, 14, 72, 420, 3, 72, 3, 72, 3, 410, 2, 73, 3, 3, 5, 4, 7, 
	5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 
	29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 
	2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 8, 
	71, 9, 73, 10, 75, 11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 17, 
	89, 18, 91, 19, 93, 20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 
	26, 107, 27, 109, 28, 111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 
	34, 123, 35, 125, 36, 127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 
	42, 139, 43, 141, 44, 143, 45, 3, 2, 33, 3, 2, 50, 59, 4, 2, 67, 67, 99, 
	99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 
	102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 
	105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 
	108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 
	111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 
	114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 
	117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 
	120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 
	123, 4, 2, 92, 92, 124, 124, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 
	67, 92, 97, 97, 99, 124, 4, 2, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 
	34, 2, 416, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 
	73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 
	2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 
	2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 
	2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 
	3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 
	2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 
	2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 
	125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 
	2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 
	3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 3, 145, 3, 2, 2, 2, 
	5, 150, 3, 2, 2, 2, 7, 153, 3, 2, 2, 2, 9, 158, 3, 2, 2, 2, 11, 160, 3, 
	2, 2, 2, 13, 166, 3, 2, 2, 2, 15, 168, 3, 2, 2, 2, 17, 170, 3, 2, 2, 2, 
	19, 172, 3, 2, 2, 2, 21, 174, 3, 2, 2, 2, 23, 176, 3, 2, 2, 2, 25, 178, 
	3, 2, 2, 2, 27, 180, 3, 2, 2, 2, 29, 182, 3, 2, 2, 2, 31, 184, 3, 2, 2, 
	2, 33, 186, 3, 2, 2, 2, 35, 188, 3, 2, 2, 2, 37, 190, 3, 2, 2, 2, 39, 192, 
	3, 2, 2, 2, 41, 194, 3, 2, 2, 2, 43, 196, 3, 2, 2, 2, 45, 198, 3, 2, 2, 
	2, 47, 200, 3, 2, 2, 2, 49, 202, 3, 2, 2, 2, 51, 204, 3, 2, 2, 2, 53, 206, 
	3, 2, 2, 2, 55, 208, 3, 2, 2, 2, 57, 210, 3, 2, 2, 2, 59, 212, 3, 2, 2, 
	2, 61, 214, 3, 2, 2, 2, 63, 216, 3, 2, 2, 2, 65, 218, 3, 2, 2, 2, 67, 220, 
	3, 2, 2, 2, 69, 229, 3, 2, 2, 2, 71, 233, 3, 2, 2, 2, 73, 238, 3, 2, 2, 
	2, 75, 241, 3, 2, 2, 2, 77, 244, 3, 2, 2, 2, 79, 249, 3, 2, 2, 2, 81, 255, 
	3, 2, 2, 2, 83, 260, 3, 2, 2, 2, 85, 269, 3, 2, 2, 2, 87, 275, 3, 2, 2, 
	2, 89, 280, 3, 2, 2, 2, 91, 291, 3, 2, 2, 2, 93, 295, 3, 2, 2, 2, 95, 297, 
	3, 2, 2, 2, 97, 299, 3, 2, 2, 2, 99, 301, 3, 2, 2, 2, 101, 303, 3, 2, 2, 
	2, 103, 306, 3, 2, 2, 2, 105, 308, 3, 2, 2, 2, 107, 310, 3, 2, 2, 2, 109, 
	313, 3, 2, 2, 2, 111, 316, 3, 2, 2, 2, 113, 319, 3, 2, 2, 2, 115, 321, 
	3, 2, 2, 2, 117, 324, 3, 2, 2, 2, 119, 326, 3, 2, 2, 2, 121, 328, 3, 2, 
	2, 2, 123, 330, 3, 2, 2, 2, 125, 332, 3, 2, 2, 2, 127, 334, 3, 2, 2, 2, 
	129, 336, 3, 2, 2, 2, 131, 338, 3, 2, 2, 2, 133, 340, 3, 2, 2, 2, 135, 
	342, 3, 2, 2, 2, 137, 355, 3, 2, 2, 2, 139, 402, 3, 2, 2, 2, 141, 404, 
	3, 2, 2, 2, 143, 418, 3, 2, 2, 2, 145, 146, 7, 101, 2, 2, 146, 147, 7, 
	113, 2, 2, 147, 148, 7, 112, 2, 2, 148, 149, 7, 101, 2, 2, 149, 4, 3, 2, 
	2, 2, 150, 151, 7, 107, 2, 2, 151, 152, 7, 104, 2, 2, 152, 6, 3, 2, 2, 
	2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 110, 2, 2, 155, 156, 7, 117, 2, 
	2, 156, 157, 7, 103, 2, 2, 157, 8, 3, 2, 2, 2, 158, 159, 7, 46, 2, 2, 159, 
	10, 3, 2, 2, 2, 160, 161, 7, 66, 2, 2, 161, 162, 7, 112, 2, 2, 162, 163, 
	7, 99, 2, 2, 163, 164, 7, 111, 2, 2, 164, 165, 7, 103, 2, 2, 165, 12, 3, 
	2, 2, 2, 166, 167, 9, 2, 2, 2, 167, 14, 3, 2, 2, 2, 168, 169, 9, 3, 2, 
	2, 169, 16, 3, 2, 2, 2, 170, 171, 9, 4, 2, 2, 171, 18, 3, 2, 2, 2, 172, 
	173, 9, 5, 2, 2, 173, 20, 3, 2, 2, 2, 174, 175, 9, 6, 2, 2, 175, 22, 3, 
	2, 2, 2, 176, 177, 9, 7, 2, 2, 177, 24, 3, 2, 2, 2, 178, 179, 9, 8, 2, 
	2, 179, 26, 3, 2, 2, 2, 180, 181, 9, 9, 2, 2, 181, 28, 3, 2, 2, 2, 182, 
	183, 9, 10, 2, 2, 183, 30, 3, 2, 2, 2, 184, 185, 9, 11, 2, 2, 185, 32, 
	3, 2, 2, 2, 186, 187, 9, 12, 2, 2, 187, 34, 3, 2, 2, 2, 188, 189, 9, 13, 
	2, 2, 189, 36, 3, 2, 2, 2, 190, 191, 9, 14, 2, 2, 191, 38, 3, 2, 2, 2, 
	192, 193, 9, 15, 2, 2, 193, 40, 3, 2, 2, 2, 194, 195, 9, 16, 2, 2, 195, 
	42, 3, 2, 2, 2, 196, 197, 9, 17, 2, 2, 197, 44, 3, 2, 2, 2, 198, 199, 9, 
	18, 2, 2, 199, 46, 3, 2, 2, 2, 200, 201, 9, 19, 2, 2, 201, 48, 3, 2, 2, 
	2, 202, 203, 9, 20, 2, 2, 203, 50, 3, 2, 2, 2, 204, 205, 9, 21, 2, 2, 205, 
	52, 3, 2, 2, 2, 206, 207, 9, 22, 2, 2, 207, 54, 3, 2, 2, 2, 208, 209, 9, 
	23, 2, 2, 209, 56, 3, 2, 2, 2, 210, 211, 9, 24, 2, 2, 211, 58, 3, 2, 2, 
	2, 212, 213, 9, 25, 2, 2, 213, 60, 3, 2, 2, 2, 214, 215, 9, 26, 2, 2, 215, 
	62, 3, 2, 2, 2, 216, 217, 9, 27, 2, 2, 217, 64, 3, 2, 2, 2, 218, 219, 9, 
	28, 2, 2, 219, 66, 3, 2, 2, 2, 220, 222, 7, 71, 2, 2, 221, 223, 7, 47, 
	2, 2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 225, 3, 2, 2, 2, 
	224, 226, 5, 13, 7, 2, 225, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 
	225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 68, 3, 2, 2, 2, 229, 230, 5, 
	41, 21, 2, 230, 231, 5, 31, 16, 2, 231, 232, 5, 37, 19, 2, 232, 70, 3, 
	2, 2, 2, 233, 234, 5, 49, 25, 2, 234, 235, 5, 55, 28, 2, 235, 236, 5, 37, 
	19, 2, 236, 237, 5, 23, 12, 2, 237, 72, 3, 2, 2, 2, 238, 239, 7, 40, 2, 
	2, 239, 240, 7, 40, 2, 2, 240, 74, 3, 2, 2, 2, 241, 242, 7, 126, 2, 2, 
	242, 243, 7, 126, 2, 2, 243, 76, 3, 2, 2, 2, 244, 245, 5, 53, 27, 2, 245, 
	246, 5, 49, 25, 2, 246, 247, 5, 55, 28, 2, 247, 248, 5, 23, 12, 2, 248, 
	78, 3, 2, 2, 2, 249, 250, 5, 25, 13, 2, 250, 251, 5, 15, 8, 2, 251, 252, 
	5, 37, 19, 2, 252, 253, 5, 51, 26, 2, 253, 254, 5, 23, 12, 2, 254, 80, 
	3, 2, 2, 2, 255, 256, 5, 41, 21, 2, 256, 257, 5, 55, 28, 2, 257, 258, 5, 
	37, 19, 2, 258, 259, 5, 37, 19, 2, 259, 82, 3, 2, 2, 2, 260, 261, 5, 51, 
	26, 2, 261, 262, 5, 15, 8, 2, 262, 263, 5, 37, 19, 2, 263, 264, 5, 31, 
	16, 2, 264, 265, 5, 23, 12, 2, 265, 266, 5, 41, 21, 2, 266, 267, 5, 19, 
	10, 2, 267, 268, 5, 23, 12, 2, 268, 84, 3, 2, 2, 2, 269, 270, 5, 17, 9, 
	2, 270, 271, 5, 23, 12, 2, 271, 272, 5, 27, 14, 2, 272, 273, 5, 31, 16, 
	2, 273, 274, 5, 41, 21, 2, 274, 86, 3, 2, 2, 2, 275, 276, 5, 23, 12, 2, 
	276, 277, 5, 41, 21, 2, 277, 278, 5, 21, 11, 2, 278, 88, 3, 2, 2, 2, 279, 
	281, 9, 29, 2, 2, 280, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 280, 
	3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 287, 3, 2, 2, 2, 284, 286, 9, 30, 
	2, 2, 285, 284, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 
	287, 288, 3, 2, 2, 2, 288, 90, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 290, 292, 
	4, 50, 59, 2, 291, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 291, 3, 
	2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 92, 3, 2, 2, 2, 295, 296, 7, 45, 2, 
	2, 296, 94, 3, 2, 2, 2, 297, 298, 7, 47, 2, 2, 298, 96, 3, 2, 2, 2, 299, 
	300, 7, 49, 2, 2, 300, 98, 3, 2, 2, 2, 301, 302, 7, 44, 2, 2, 302, 100, 
	3, 2, 2, 2, 303, 304, 7, 63, 2, 2, 304, 305, 7, 63, 2, 2, 305, 102, 3, 
	2, 2, 2, 306, 307, 7, 64, 2, 2, 307, 104, 3, 2, 2, 2, 308, 309, 7, 62, 
	2, 2, 309, 106, 3, 2, 2, 2, 310, 311, 7, 64, 2, 2, 311, 312, 7, 63, 2, 
	2, 312, 108, 3, 2, 2, 2, 313, 314, 7, 62, 2, 2, 314, 315, 7, 63, 2, 2, 
	315, 110, 3, 2, 2, 2, 316, 317, 7, 35, 2, 2, 317, 318, 7, 63, 2, 2, 318, 
	112, 3, 2, 2, 2, 319, 320, 7, 35, 2, 2, 320, 114, 3, 2, 2, 2, 321, 322, 
	7, 60, 2, 2, 322, 323, 7, 63, 2, 2, 323, 116, 3, 2, 2, 2, 324, 325, 7, 
	63, 2, 2, 325, 118, 3, 2, 2, 2, 326, 327, 7, 93, 2, 2, 327, 120, 3, 2, 
	2, 2, 328, 329, 7, 95, 2, 2, 329, 122, 3, 2, 2, 2, 330, 331, 7, 61, 2, 
	2, 331, 124, 3, 2, 2, 2, 332, 333, 7, 125, 2, 2, 333, 126, 3, 2, 2, 2, 
	334, 335, 7, 127, 2, 2, 335, 128, 3, 2, 2, 2, 336, 337, 7, 42, 2, 2, 337, 
	130, 3, 2, 2, 2, 338, 339, 7, 43, 2, 2, 339, 132, 3, 2, 2, 2, 340, 341, 
	7, 48, 2, 2, 341, 134, 3, 2, 2, 2, 342, 350, 7, 36, 2, 2, 343, 344, 7, 
	94, 2, 2, 344, 349, 11, 2, 2, 2, 345, 346, 7, 36, 2, 2, 346, 349, 7, 36, 
	2, 2, 347, 349, 10, 31, 2, 2, 348, 343, 3, 2, 2, 2, 348, 345, 3, 2, 2, 
	2, 348, 347, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 350, 
	351, 3, 2, 2, 2, 351, 353, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 353, 354, 
	7, 36, 2, 2, 354, 136, 3, 2, 2, 2, 355, 356, 5, 89, 45, 2, 356, 357, 5, 
	133, 67, 2, 357, 358, 5, 89, 45, 2, 358, 138, 3, 2, 2, 2, 359, 361, 5, 
	13, 7, 2, 360, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 360, 3, 2, 2, 
	2, 362, 363, 3, 2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 360, 3, 2, 2, 2, 364, 
	365, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 368, 7, 48, 2, 2, 367, 369, 
	5, 13, 7, 2, 368, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368, 3, 2, 
	2, 2, 370, 371, 3, 2, 2, 2, 371, 403, 3, 2, 2, 2, 372, 374, 5, 13, 7, 2, 
	373, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 373, 3, 2, 2, 2, 375, 
	376, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 7, 48, 2, 2, 378, 379, 
	5, 67, 34, 2, 379, 403, 3, 2, 2, 2, 380, 382, 5, 13, 7, 2, 381, 380, 3, 
	2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 383, 384, 3, 2, 2, 
	2, 384, 386, 3, 2, 2, 2, 385, 381, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 
	387, 3, 2, 2, 2, 387, 389, 7, 48, 2, 2, 388, 390, 5, 13, 7, 2, 389, 388, 
	3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 389, 3, 2, 2, 2, 391, 392, 3, 2, 
	2, 2, 392, 393, 3, 2, 2, 2, 393, 394, 5, 67, 34, 2, 394, 403, 3, 2, 2, 
	2, 395, 397, 5, 13, 7, 2, 396, 395, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 
	396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 
	5, 67, 34, 2, 401, 403, 3, 2, 2, 2, 402, 364, 3, 2, 2, 2, 402, 373, 3, 
	2, 2, 2, 402, 385, 3, 2, 2, 2, 402, 396, 3, 2, 2, 2, 403, 140, 3, 2, 2, 
	2, 404, 405, 7, 49, 2, 2, 405, 406, 7, 49, 2, 2, 406, 410, 3, 2, 2, 2, 
	407, 409, 11, 2, 2, 2, 408, 407, 3, 2, 2, 2, 409, 412, 3, 2, 2, 2, 410, 
	411, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411, 413, 3, 2, 2, 2, 412, 410, 
	3, 2, 2, 2, 413, 414, 7, 12, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 8, 71, 
	2, 2, 416, 142, 3, 2, 2, 2, 417, 419, 9, 32, 2, 2, 418, 417, 3, 2, 2, 2, 
	419, 420, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 
	422, 3, 2, 2, 2, 422, 423, 8, 72, 2, 2, 423, 144, 3, 2, 2, 2, 22, 2, 222, 
	227, 282, 285, 287, 293, 348, 350, 362, 364, 370, 375, 383, 385, 391, 398, 
	402, 410, 420, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'conc'", "'if'", "'else'", "','", "'@name'", "", "", "'&&'", "'||'", 
	"", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", 
	"'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'['", "']'", "';'", 
	"'{'", "'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", 
	"SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", 
	"MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", 
	"SET", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", 
	"RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", 
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C", "D", 
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", 
	"T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "NIL", "RULE", 
	"AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", 
	"SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", 
	"GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "LSQARE", "RSQARE", 
	"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", 
	"DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewgengineLexer(input antlr.CharStream) *gengineLexer {

	l := new(gengineLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0 = 1
	gengineLexerT__1 = 2
	gengineLexerT__2 = 3
	gengineLexerT__3 = 4
	gengineLexerT__4 = 5
	gengineLexerNIL = 6
	gengineLexerRULE = 7
	gengineLexerAND = 8
	gengineLexerOR = 9
	gengineLexerTRUE = 10
	gengineLexerFALSE = 11
	gengineLexerNULL_LITERAL = 12
	gengineLexerSALIENCE = 13
	gengineLexerBEGIN = 14
	gengineLexerEND = 15
	gengineLexerSIMPLENAME = 16
	gengineLexerINT = 17
	gengineLexerPLUS = 18
	gengineLexerMINUS = 19
	gengineLexerDIV = 20
	gengineLexerMUL = 21
	gengineLexerEQUALS = 22
	gengineLexerGT = 23
	gengineLexerLT = 24
	gengineLexerGTE = 25
	gengineLexerLTE = 26
	gengineLexerNOTEQUALS = 27
	gengineLexerNOT = 28
	gengineLexerASSIGN = 29
	gengineLexerSET = 30
	gengineLexerLSQARE = 31
	gengineLexerRSQARE = 32
	gengineLexerSEMICOLON = 33
	gengineLexerLR_BRACE = 34
	gengineLexerRR_BRACE = 35
	gengineLexerLR_BRACKET = 36
	gengineLexerRR_BRACKET = 37
	gengineLexerDOT = 38
	gengineLexerDQUOTA_STRING = 39
	gengineLexerDOTTEDNAME = 40
	gengineLexerREAL_LITERAL = 41
	gengineLexerSL_COMMENT = 42
	gengineLexerWS = 43
)

