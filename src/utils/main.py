import os
import sys
import argparse
import logging
from datetime import datetime

def setup_logging():
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler('devops-scripts.log'),
            logging.StreamHandler(sys.stdout)
        ]
    )

def parse_args():
    parser = argparse.ArgumentParser(description='DevOps Scripts CLI Tool')
    parser.add_argument('-c', '--config', type=str, required=True, help='Path to configuration file')
    parser.add_argument('-o', '--output', type=str, default='output.txt', help='Output file path')
    return parser.parse_args()

def load_config(config_path):
    if not os.path.exists(config_path):
        raise FileNotFoundError(f"Config file not found: {config_path}")
    with open(config_path, 'r') as file:
        return file.read()

def process_data(config_data, output_path):
    timestamp = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    processed_data = f"Processed on {timestamp}:\n{config_data}"
    with open(output_path, 'w') as file:
        file.write(processed_data)
    logging.info(f"Data processed and saved to {output_path}")

def main():
    setup_logging()
    args = parse_args()
    try:
        config_data = load_config(args.config)
        process_data(config_data, args.output)
    except Exception as e:
        logging.error(f"An error occurred: {e}")
        sys.exit(1)

if __name__ == '__main__':
    main()