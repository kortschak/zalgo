zalgo finally brings support for Zalgo text to Go.

    package zalgo_test

    import (
        "fmt"
        "github.com/kortschak/zalgo"
        "os"
    )

    const invoke = `To invoke the hive-mind representing chaos.
    Invoking the feeling of chaos.
    With out order.
    The Nezperdian hive-mind of chaos. Zalgo.
    He who Waits Behind The Wall.
    ZALGO!
    `

    func Example_1() {
        z := zalgo.NewCorrupter(os.Stdout)

        z.Zalgo = func(n int, z *zalgo.Corrupter) {
            z.Up += 0.1
            z.Middle += complex(0.01, 0.01)
            z.Down += complex(real(z.Down)*0.1, 0)
        }

        z.Up = complex(0, 0.2)
        z.Middle = complex(0, 0.2)
        z.Down = complex(0.001, 0.3)

        fmt.Fprintln(z, invoke)

        // Output:
        // Eternal happiness.
    }


`got:`

To ͪ͞i͝n͙v͟o̟ke̷ ̢̝t͓ͪh̨e̖ ͨ͟h͇iv̋e-̪m̘i̹n̑d̝ ͏̺re̞̓͡p͢r͡és̷e̎͞ņ̿͐͒tį̊ng̃͢ ̓̽c̨̰ͬ̇ḫ͞a͢o͍͑s̐ͣ́.

I̛ͭ̌n̪̉̓̀v̒͢ǫk͐͝in̸͆g ͚̑̔͠t͓̽ͤ͡h̦͒͆͜e̜̿ ͢f̸e͂͏̻e̵ļi͂͗̀n̡ͤ̋g̢̾ ̵̹́̊ơf̲͞ ̉͢c̎̏̉͟h̘̒̾̆͠â̧ǫ͖̈̂̓ͫs̫̳.̋҉̪

W͍̗͟í̗t̺̬̏͊̌̕ẖ̢̎̏ ͔̈͝ŏ̻̊ͧ̀u̕t̨̲ͯ̿ ̭͂ͨͮ͆ͮ̈͢o̻̍͡r̠̎͢ͅḏ͐͟ë̢͙̇̿̽r͏͓̮.̴̋

T͇͇͌ͣ͊͝h̦̳̽ͫͬ́e̥̤̬̐͛ͥ͠ ̶̱̝͎͇ͦ̈́͑͊N̨̹͎ͤȩ͈̻̲̕ͅz̨̫̫̼͕̱̜̆͢p̵̟͙̤͇̈́͂̎̽̀ė̷ͦ͏͚̤̘̜͓̲r̶͎͑̄̀d̛͖̠̣̣̦̽͝į̷͔̲̦̩̜̮͈͒ą̨͈̘̥̥͉ͨ͑͐n̸̲̜̳͓̱̙͌͡ ͋̈̍̑͝҉̥͈͎̝h̢̥̹̣̖̱̓͊͡ͅiͦ̋̄͞҉̜͎̳̗ͅv̵̨̙͙̰͍̼ͧ̈́ͅę̣͓͈̣̱̞̚͠ͅ-̷̛͖̜͎̠̼̻͑͆m͙̺̦̱͚̋ͣ͘͞ͅi̷̷̲̯̣̺͓͉̰͈n͚̘̣̳̥̥̂̊͢͞d̨̘̣͓̮̱̓͌́̕ ̤̝̼͇͇̩̘̳͘͜ǫ̴̣̙͙̦̝̠̮̈f̡͙̞̬̤̪͙̊͛͟ ͕̖̉͒ͨ̒ͨ́͝ͅċ̡͈͔̦̙̤͎͞ͅh͐̐҉̲̪̼̮̫̼͜ą̰̘̙͙͓̘̲̐͠o̡̩͉͓̫̼͗͂́͟ṡ̸̪͎̱̭͙͎ͭ͜.̝̗̪̑ͭ̿͆ͩ̀͞ ͌͐̅ͦ̋̍͛̌͏̨Z̝̝͎̫ͭ͊ͤ̚͟͝ą̳͙̹͉͖͉̎͘ͅl̸͈͖̝̣ͩͬͣ̚͝g̸̷̜̝͍̲͛̊̎͊ó̵̗͉̠̯̬͎͐͝.͡͏̩̥̖͉͓͙̺̲

H̢͍̺͖͂̓͗ͬ̍̀e̼͇̦̰͔͖͂͒̕͡ ̧̭͈̖̟͕̃̓ͦ͠w̴̷͈̰̻͈͕̋͒̓h̑̔̀̿ͤ̓͏̱̰͘o̵̸͈͔̫͔͍̤̓̊ ̴͖̞͚̯̝͐͛ͫ͡W̷̱͖̲̤̬̃̔͂͝ạ̸͍̇̔̑̑͆ͩ͝ị̢̢̯̠͕̞̓ͪ͛ţ̴̜̝̞̭̑ͬ̄ͥs̛̲͙͕̪̓̿̑͛͡ ̣̰͎̽̏̄́̊̀͘B̨͔͕̮̝͕̩̓̂͡ẹ̴͍̟̮̅͛ͫ̕ͅh̆̅ͨ̒ͨ͏̮̙̠͠í̻̩̭̫͎ͯ̚͡͠ṅ̵̢̜̪̮̩̣ͤ͂d̘̮͇̟̺̖ͦ͒̀͜ ̸̰̣̳̭̜̈́́̀ͅT̸̢͕̬̯͎͒ͤ͗ͦh̷̵̥̫̻͓͔̖̃̄ę̞͇̗̣̫̟̼̎͡ ̴̧͎̠͇̥̪̬ͮ͛W̶̨͖͖͈͎̌͆̌͐a̷̧̺̦͚̯̹ͣ̚ͅl̢̥̹̜͑̑ͣͩ̂͟l̃ͧͩͭ͏̹͈̬͡ͅ.̢̯̔́͆ͩͪ͊̑͟

Z̛̝͕̫͎ͦ̓̓́͡Ḁ̧͖̣͎̋̀ͯ̈̕L̷̹̥̣̪̮̩̅̄͝G̱̭͔̜̹͊͆̃͝͠O̍͐͐͏̷̟̰̼͉̰!̴͎̩̣̇͗̈́̀̅͡



`want:`

Eternal happiness.
