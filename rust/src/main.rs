extern crate rand;

use std::env;
use std::process;

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
    let mut args = env::args();
    let mut o = DEFAULT_OPTIONS.clone();
    let mut n: u16 = 1;
    // first argument is program name, throw it away
    args.next();
    loop {
        match args.next() {
            Some(a) => {
                match a.as_str() {
                    "-xu" => o.upper = false,
                    "-xl" => o.lower = false,
                    "-xn" => o.number = false,
                    "-xs" => o.symbol = false,
                    "-l" => {
                        match args.next() {
                            Some(l) => {
                                match l.parse::<u16>() {
                                    Ok(i) => o.length = i,
                                    _ => {
                                        eprintln!("Failed to parse length argument as integer");
                                        process::exit(1);
                                    },
                                }
                            },
                            _ => {
                                eprintln!("Length argument given but no value");
                                process::exit(1);
                            },
                        }
                    },
                    "-n" => {
                        match args.next() {
                            Some(i) => {
                                match i.parse::<u16>() {
                                    Ok(i) => n = i,
                                    _ => {
                                        eprintln!("Failed to parse number argument as integer");
                                        process::exit(1);
                                    },
                                }
                            },
                            _ => {
                                eprintln!("Number argument given but no value");
                                process::exit(1);
                            },
                        }
                    },
                    "-version" => {
                        println!(env!("CARGO_PKG_VERSION"));
                        process::exit(1);
                    },
                    _ => eprintln!("Unknown argument: {}", a),
                }
            },
            None => break,
        }
    }
    let g = new(o);
    for _ in 0..n {
        let s = g.generate();
        println!("{}", s);
    }
}
