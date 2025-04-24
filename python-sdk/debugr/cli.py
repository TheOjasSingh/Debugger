import click
from debugr.sdk import DebugrClient

@click.group()
def cli():
    pass

@cli.command()
@click.argument('pid', type=int)
@click.option('--url', default='http://localhost:8080', help='Debugr base URL')
def trace(pid, url):
    """Fetch stack trace for a process by PID."""
    client = DebugrClient(url)
    try:
        stack = client.trace(pid)
        click.echo(stack)
    except Exception as e:
        click.echo(f"Error: {e}", err=True)

if __name__ == "__main__":
    cli()
