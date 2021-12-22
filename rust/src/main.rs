extern crate rand;

use std::env;

const MAX_TRIES: u8 = 10;

const UPPER: &str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
const LOWER: &str = "abcdefghijklmnopqrstuvwxyz";
const NUMBER: &str = "1234567890";
const SYMBOL: &str = "`~!@#$%^&*()_+-=[]{}\\|;:'\",./<>?";

#[derive(Copy, Clone)]
struct Options {
    upper: bool,
    lower: bool,
    number: bool,
    symbol: bool,
    length: u16,
}

struct Generator {
    opts: Options,
    charset: String,
}

static DEFAULT_OPTIONS: Options = Options{
    upper: true,
    lower: true,
    number: true,
    symbol: true,
    length: 16,
};

fn new(o: Options) -> Generator {
    let mut g = Generator{
        opts: o,
        charset: String::new(),
    };
    if g.opts.upper {
        g.charset.push_str(UPPER);
    }
    if g.opts.lower {
        g.charset.push_str(LOWER);
    }
    if g.opts.number {
        g.charset.push_str(NUMBER);
    }
    if g.opts.symbol {
        g.charset.push_str(SYMBOL);
    }
    return g;
}

impl Generator {
    fn generate(&self) -> String {
        let mut tries = 0;
        let mut s = String::new();

        loop {
            tries += 1;
            if tries > MAX_TRIES {
                panic!("Circuit breaker tripped. Failed to generate string in {} tries.", MAX_TRIES);
            }
            s.clear();
            let mut x: usize;
            let mut y: usize;
            for _ in 0..self.opts.length {
                x = rand::random();
                y = x % self.charset.len();
                s.push(self.charset.chars().nth(y).unwrap());
            }
            if self.opts.upper && !UPPER.chars().any(|c| s.contains(c) ) {
                continue
            }
            if self.opts.lower && !LOWER.chars().any(|c| s.contains(c) ) {
                continue
            }
            if self.opts.number && !NUMBER.chars().any(|c| s.contains(c) ) {
                continue
            }
            if self.opts.symbol && !SYMBOL.chars().any(|c| s.contains(c) ) {
                continue
            }
            break
        }
        return s;
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.contains(&String::from("-version")) {
        println!("{}", env!("CARGO_PKG_VERSION"));
        return;
    }
    let mut o = DEFAULT_OPTIONS.clone();
    o.upper = !args.contains(&String::from("-xu"));
    o.lower = !args.contains(&String::from("-xl"));
    o.number = !args.contains(&String::from("-xn"));
    o.symbol = !args.contains(&String::from("-xs"));
    match args.iter().position(|x| x == &String::from("-l")) {
        Some(y) => {
            match args.get(y + 1) {
                Some(z) => {
                    match z.parse::<u16>() {
                        Ok(j) => o.length = j,
                        _ => (/*TODO Error, -l arg given, but following value failed to parse to int*/),
                    }
                },
                _ => (/*TODO Error, -l arg given but no value following */),
            }
        },
        _ => (/* no length argument, this is ok */),
    }
    let g = new(o);
    let s = g.generate();

    println!("{}", s);
}
