using System;
using System.Collections.Generic;
using System.IO;
using System.Runtime.CompilerServices;
using ApiFarm.Models.Impl;
using Newtonsoft.Json;

[assembly: InternalsVisibleTo("ApiFarm.Test")]

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Loads sample video games from data file.
    /// </summary>
    public class VideoGameDataLoader : IDataLoader<VideoGame>
    {
        private Func<string, string> readFile = path => File.ReadAllText(path);
        private Func<string, IEnumerable<VideoGame>> parseJson = text => JsonConvert.DeserializeObject<List<VideoGame>>(text);

        /// <summary>
        /// Initializes a new instance of the <see cref="VideoGameDataLoader"/> class.
        /// </summary>
        public VideoGameDataLoader()
        {
        }

        internal VideoGameDataLoader(Func<string, string> readFile, Func<string, IEnumerable<VideoGame>> parseJson)
        {
            this.readFile = readFile;
            this.parseJson = parseJson;
        }

        /// <summary>
        /// Reads and deserializes json data to <see cref="VideoGame"/> instances.
        /// </summary>
        /// <param name="dataPath">The path to the data sample file.</param>
        /// <returns><see cref="VideoGame"/> instances from data samples.</returns>
        public IEnumerable<VideoGame> Load(string dataPath)
        {
            var jsonSampleData = this.readFile(dataPath);
            return this.parseJson(jsonSampleData);
        }
    }
}
