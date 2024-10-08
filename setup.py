from setuptools import setup, find_packages

setup(
    name="blitz_env",
    version="0.1.0",
    packages=find_packages(),  # Automatically find packages in subfolders
    author="Mitch Chris",
    author_email="mitchchris@example.com",
    description="Used for bot blitz",
    long_description="",
    long_description_content_type='text/markdown',
    url="https://github.com/yourusername/your-repo",
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    install_requires=[
        "pandas==2.1.4",
        "nfl-data-py==0.3.1",
        'protobuf==5.27.4',
        'matplotlib==3.9.2',
        'beautifulsoup4==4.12.3',
        'requests==2.32.3'
    ],
    package_data={
        'blitz_env': ['player_ranks_2024.csv', 'player_ranks_2023.csv', 'player_ranks_2022.csv', 'player_ranks_2021.csv'],
    },
    python_requires='>=3.6',
)
