function reverse(input: string): string {
  if (input.length === 0) {
    return ""
  }
  // all the chars in the string except the first and add it to the end!
  return reverse(input.slice(1)) + input.at(0)
}

console.log(reverse("abc"))
