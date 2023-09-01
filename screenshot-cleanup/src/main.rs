use std::fs::metadata;
use std::path::{Path, PathBuf};

fn print_type<T>(_: &T) {
    println!("{}", std::any::type_name::<T>());
}

fn main() {
    let path = Path::new("/Users/LJPurcell/Desktop/");
    let cd_result = std::env::set_current_dir(path);

    let entries_result = match cd_result {
        Ok(_) => std::fs::read_dir("."),
        Err(_) => panic!("Couldn't change directory using path: {}", path.display()),
    };

    let entries = match entries_result {
        Ok(entries) => entries,
        Err(_) => panic!("Couldn't read current directory at: {}", path.display()),
    };

    for entry in entries {
        let file = entry.unwrap().path();

        if file.is_file() {
            if file.extension().is_some() {
                let extension = file.extension().unwrap();

                if extension == "png" {
                    println!("Found the png");
                }

                let file_name = file.to_str();
                println!("{:?}", file_name);
            };
        }
    }
}
