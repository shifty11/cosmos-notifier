fn main() -> Result<(), Box<dyn std::error::Error>> {
    let proto_files: Vec<std::path::PathBuf> = std::fs::read_dir("api")?
        .map(|res| res.map(|e| e.path()))
        .collect::<Result<Vec<_>, std::io::Error>>()?;
    let proto_files_refs: Vec<&std::path::Path> = proto_files
        .iter()
        .map(|p| p.as_path())
        .filter(|p| p.extension().unwrap_or_default() == "proto")
        .collect();

    tonic_build::configure()
        .build_server(false)
        .compile(&proto_files_refs, &["api"])?;
    Ok(())
}
