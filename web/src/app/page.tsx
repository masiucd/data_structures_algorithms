import {Button, Flex, ThemePanel} from "@radix-ui/themes";
import {revalidatePath} from "next/cache";

let themePanelOpen = false;

async function toggleThemePanel() {
  "use server";
  themePanelOpen = !themePanelOpen;
  revalidatePath("/");
}

export default function Home() {
  let isProduction = process.env.NODE_ENV === "production";
  return (
    <>
      <h1>Hi</h1>
      <p>
        Lorem ipsum dolor sit, amet consectetur adipisicing elit. Molestias
        adipisci molestiae odio numquam nisi iste porro totam debitis, eos at
        voluptatum libero velit nemo obcaecati accusantium quasi, a tempora
        quam.
      </p>
      <Flex gap="3">
        <Button className="cursor-pointer" variant="soft">
          Algorithms
        </Button>
        <Button variant="soft">Data Structures</Button>
      </Flex>
      {!isProduction && (
        <form action={toggleThemePanel}>
          <Button variant="classic" title="Toggle theme panel" type="submit">
            view theme panel
          </Button>
        </form>
      )}
      {!isProduction && themePanelOpen && <ThemePanel />}
    </>
  );
}
