use std::io::Write;
use std::process::Command;

use anyhow::Result;
use flowe::json;
use tempfile::NamedTempFile;

fn main() -> Result<()> {
    let args = std::env::args().collect::<Vec<String>>();
    let input_file = &args[1];
    let exec_step = &args[2];
    let input_content = std::fs::read_to_string(input_file)?;
    let plan: json::PlanJson = serde_json::from_str(&input_content)?;
    do_step(&plan, exec_step)?;
    Ok(())
}

fn find_step<'a>(plan: &'a json::PlanJson, step_name: &str) -> Option<&'a json::StepJson> {
    plan.steps.iter().find(|step| step.name == step_name)
}

fn do_step(plan: &json::PlanJson, step_name: &str) -> Result<()> {
    let step = find_step(plan, step_name);
    if let Some(step) = step {
        let step_exec = step.exec.clone();
        let mut temp_file = NamedTempFile::new()?;
        temp_file.write_all(step_exec.as_bytes())?;
        temp_file.flush()?;
        let path = format!("{}", temp_file.path().display());
        let op = Command::new("/bin/bash")
            .arg(path)
            .output()?;
        if op.status.code() != Some(0) {
            return Err(anyhow::anyhow!("step {} failed", step_name));
        }
        println!("{}", String::from_utf8_lossy(&op.stdout));
    } else {
        return Err(anyhow::anyhow!("step {} not found", step_name));
    }
    Ok(())
}
