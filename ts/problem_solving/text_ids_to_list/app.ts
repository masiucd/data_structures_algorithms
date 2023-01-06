const __dirname = new URL(".", import.meta.url).pathname

const file = await Deno.readTextFile(__dirname + "text.txt")

const lines = file.split("\n")
