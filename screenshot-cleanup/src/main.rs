use std::env;
use std::{fs, path::Path};

fn main() {
    let path = Path::new("/Users/LJPurcell/Desktop/");
    let cd_result = env::set_current_dir(path);
    let file_paths = match cd_result {
        Ok(cd_result) => fs::read_dir(path),
        Err(error) => panic!("Couldn't change to Desktop using path: {}", path.display()),
    };

    for path in file_paths.unwrap() {
        let fp = path.unwrap();
        println!("{}", fp.path().display());
    }
}
