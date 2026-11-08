package highlighter

const goSample = `
package main

import "fmt"

// Greet says hello to name, or to the world.
func Greet(name string) string {
	if name == "" {
		return "hello, world"
	}
	return fmt.Sprintf("hello, %s", name)
}
`

const tsSample = `
export const fetchMessages = async (id: string): Promise<Message[]> => {
  const res = await fetch(` + "`/api/chat/${id}`" + `);
  if (!res.ok) throw new Error(res.statusText);
  return (await res.json()) as Message[];
};
`
