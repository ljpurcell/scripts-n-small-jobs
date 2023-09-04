use std::{fs, path::Path};

fn main() {
    let path = Path::new("/Users/LJPurcell/Desktop/");
    let cd_result = std::env::set_current_dir(path);

    let read_dir_result = match cd_result {
        Ok(_) => std::fs::read_dir("."),
        Err(_) => panic!("Couldn't change directory using path: {}", path.display()),
    };

    let entries = match read_dir_result {
        Ok(entries) => entries,
        Err(_) => panic!("Couldn't read current directory at: {}", path.display()),
    };

    for entry in entries {
        let file_path = match entry {
            Ok(entry) => entry.path(),
            Err(_) => panic!("Couldn't get entry at: {:?}", entry),
        };

        if file_path.is_file() {
            if file_path.extension().is_some() {
                let extension = file_path
                    .extension()
                    .unwrap_or_else(|| panic!("Could't get extension for file: {:?}", file_path));

                let file_name = file_path.to_str().unwrap_or_else(|| {
                    panic!(
                        "Couldn't convert file path to string for file {:?}",
                        file_path
                    )
                });

                if extension == "png" && file_name.starts_with("./Screenshot 2023-") {
                    fs::remove_file(file_name).unwrap_or_else(|file_name| panic!("Error when removing file: {:?}", file_name));
                }
            };
        }
    }
}
