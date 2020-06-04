/**
 * tcmd-handler (examples/typescript)
 *
 * This is a simple example showing how you can use Deno (https://deno.land) with tcmd.
 * Use `deno install ./examples/typescript/tcmd-handler.ts` then `tcmd-server` to run it.
 * `tcmd localhost` can be used to test it; simply type "test" and you should receive "ok".
 */

// We need this to parse tcmd's input.
import { parse } from "https://deno.land/std/flags/mod.ts";

// Let's parse the arguments.
const args = parse(Deno.args);
// These are the arguments that tcmd passes.
const req = {
  method: args._[0],
};

// This switches on the `method` key that clients send.
switch (req.method) {
  case "test":
    // If the method's name is test, we're simply going to send back "ok".
    console.log("ok");
}
