let vowels = Object.freeze(["a", "i", "e", "o", "u"]);
export function reverseVowels(s: string) {
  let result = "";
  let vowelsInString = [];
  for (let char of s) {
    if (vowels.includes(char.toLowerCase())) {
      vowelsInString.push(char);
    }
  }

  for (let char of s) {
    if (vowels.includes(char.toLowerCase())) {
      result += vowelsInString.pop();
    } else {
      result += char;
    }
  }
  return result;
}
