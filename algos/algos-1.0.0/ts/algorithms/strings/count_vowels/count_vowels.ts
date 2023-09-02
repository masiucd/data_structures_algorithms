const vowels = ["a", "e", "i", "o", "u"];
function isVowel(char: string): boolean {
  return vowels.includes(char);
}

export function countVowels(input: string): number {
  if (input === "") return 0;
  let count = isVowel(input[0]) ? 1 : 0;
  return count + countVowels(input.substring(1));
}

export function countVowelsIterative(input: string): number {
  let count = 0;
  for (let i = 0; i < input.length; i++) {
    if (isVowel(input[i])) count++;
  }
  return count;
}

export function countVowelsFunctional(input: string): number {
  return input.split("").filter(isVowel).length;
}
