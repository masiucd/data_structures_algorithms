// Good morning! Here’s your coding interview problem for today.

// This problem was asked by Amazon.

// Run-length encoding is a fast and simple method of encoding strings. The basic idea is to represent repeated successive characters as a single count and character. For example, the string “AAAABBBCCDAA” would be encoded as “4A3B2C1D2A”.

// Implement run-length encoding and decoding. You can assume the string to be encoded have no digits and consists solely of alphabetic characters. You can assume the string to be decoded is valid.

const countChars = (input: string): string => {
  let result = ""
  let count = 1
  for (let i = 1; i <= input.length; i++) {
    const currentChar = input.charAt(i - 1)
    if (input.charAt(i) === currentChar) {
      count++
    } else {
      result = addCount(result, count, currentChar)
      count = 1
    }
  }
  return result
}

function addCount(result: string, count: number, char: string): string {
  result += count
  result += char
  return result
}

console.log(countChars("AAAABBBCCDAA")) // 4A3B2C1D2A;
