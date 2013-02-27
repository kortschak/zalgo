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

    func ExampleWrite() {
        z := zalgo.NewCorrupter(os.Stdout)

        z.Zalgo = func(n int, z *zalgo.Corrupter) {
            z.Up += 0.0001
            z.Middle += 0.001
            d := complex(real(z.Down)*0.1, 0)
            z.Down += d
        }

        z.Up = complex(0, 0.2)
        z.Middle = complex(0, 0.2)
        z.Down = complex(0.001, 0.3)

        fmt.Fprintln(z, invoke)

        // Output:
        // Eternal happiness.
    }


`got:`

To ͪ͞in̝v̼oͅkeͅ ̮t̥ͥhẹ ͈h̍͘ive̒-̭mind̀ r̄e̹p͘resͥe̳n͛ti̡͉n̸g͢ cͤͅh́aos.̯̎

Invọ̵kin̖g̼ ̧͍th̩ė ̷fe̮eḻi̱ng ̐o̶fͥ c̜hao̯s̪ͯ.̳

W̪i̠t̀h͔ ̣͐o̞u̻t͆ o̝r̢͙͚͓͍d̵̘͚ȩ̼̦r̜̿.̫̜

T͔̥h͈̫̗̞͠e̲͖ ̝̹̟̹̼N̛̝̣̲̲̠ẻ͖̲͙͎̱̠z͖͈̘͓̗̳͠p͍͎͖͕̯̫̳̂͡e̙̪̞̥̬̠̋r̪͍̗d̯̪̗̮̮̠̺̬ͅi̝̰̣̱̰͍͔̼̠͜a̧̻͕̺̫ͅn͎̲͔̯̜͚͔̘͇̭ ̦̤̠͓̝̯̱͎̦ͅh̲͙̙̺̮͓̮͖̬̿i̮͚̯̮͈͉͎̹̪ͅṿ̹͙͇͇̗̗̘̪̦e̟̻͈͙̥̤̬̼͖̥-̹͎̲͈̻̲͕̭̩ͅm̼͕̱̜͖̫͖̜͍̿i̘̜͓̲̫̱̘͎͋͠n͖̠̣̣̦͇̫̼͔̲d̻͈̘̥̥͉͕͔͐ͅ ̩͉̝̥͈͎̝̪̮̣o̤̼̭̦̘͓̜͎̳ͅf͓̱͇͔̪͕̣̫̄ͅ ͖̜͎̠̼̻̖̪̫͖c̙͚͎̗̲̯̣̺͓͉h͈̱͍̞̬͔̺̙̱͎a̩̼̤̝̼͇͇̩̘̳o̦̘̬̝̝̰̜̩̪̭s̵̯͇̠͇̯̱͕̖ͅ.̙͖͍̮͙͎͔̗̝̖ ̴̦̯͖͍̯̭̻͔̻Z̰̫̼̼̜̙͓͈̟̬a̪̗̤̣̮̟̘̠͈͞l̻̫̥̜̻̟̪̼̜͛g̮̰̺̮̝̝͎̫ͭͅo̲̮̻͕̣̠̣̪͓͕.̣͇̙̦͚͚̭̻̭ͅ

H̝̥̲̲͕͇̺̥̩͍eͨ҉̪̟̼̝̹̗͍̥ ̬̙̰̖̹̞̜̞̫ͅw̢̪͙̩̯͇͎͈̩ͅh͇͍̮͓̹̬̦͇̙͚ǫ͖̦̣͖̬̝̟̤͈ ͖̝̭̱̟̖̘͓̙ͅW̖͉̰̠̳̪̦̰̘̪a̫̹̦͙̱̺̱͖̲̤i̩̝͚̲̯̟̥͖̞͢t̙̟̩̪̼̳̞̞̚ͅs̖͇̥͇̘̟͓͍̝̱ ̙͉̲̱̘̥̳̺̥͔B̻̻͙͎̻͓̯͚͉̘ḙ̠̙͚̣͉̬̰̱ͅh̝̟͇̣̼̫̥̺̝̯i̩͍͚̲̹͚̤̘̮ͅṋ̭̜͚̣͓͕̪̖͔d̗̖̰͖̱̹̹͖̞͕ ̳̞͖̭̬̯̪̯̭̜T͚̥͕͍͚̰̭̠̥̪h͕̤͕̞̱͔̮̰ͅͅe͈̦̘̞̘͎̼̮͖̪ ̷̪̫͙̫̠̲͕̤̱W҉̻̮͙͈̺̖̥͕͇a̷̘̖͉̠̩̣̰͍̤l̻͉̜̻̣̪͕͓̩͓l̳͚͕͚̭͇͚͉͔̙.̦̠̹̻̬̤̪̥̜͇

Z̖͔͎̫͔̻͍͚̞̭A̼̬̩̪͚͎͕̯̬͉L̖͎̬̰̪͓̗̖̺ͅG̱̘̲͙̺̹͈͍̭̺O҉̖͚̙̥̥͓̝͍̪!̳̳̟͎̝̳͇̘̟͚


`want:`

`Eternal happiness.`
