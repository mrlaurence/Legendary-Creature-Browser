function Analyse() {
    this.rS = [];
    this.description = [];
    this.immediate = [];
    this.stopwords = ['in','the', 'a','i','me','my','myself','we','our','ours','ourselves','you','your','yours','yourself','yourselves','he','him','his','himself','she','her','hers','herself','it','its','itself','they','them','their','theirs','themselves','what','which','who','whom','this','that','these','those','am','is','are','was','were','be','been','being','have','has','had','having','do','does','did','doing','a','an','the','and','but','if','or','because','as','until','while','of','at','by','for','with','about','against','between','into','through','during','before','after','above','below','to','from','up','down','in','out','on','off','over','under','again','further','then','once','here','there','when','where','why','how','all','any','both','each','few','more','most','other','some','such','no','nor','not','only','own','same','so','than','too','very','s','t','can','will','just','don','should','now']
    this.nlp;
}



Analyse.prototype = {

    /*
     *sentences with singular pronouns
     *find comparable adjectives
     *find like-a nouns and hyphenated nouns
     *use copula to relate specific nouns to subject
     *adverb followed by noun pairs
     *extract relevant places/locations using .places() and demonyms
     */
    init: function() {
        this.nlp = require('compromise');
        this.nlp.extend(require('compromise-ngrams'))
        this.nlp.extend(require('compromise-sentences'))
        console.log("\n Loaded modules \n");
    },

    //subject can be a header value(will help for context of sentences!)
    sentences: function(input, subject="-") {
        //extract text from wiki paragraph
        let text = input.replace(/[^A-Za-z.]+/gi, " ");
        let sentences = this.nlp(text).sentences().text().split(".");
        sentences = sentences.filter(sentences => sentences);
        //some immediate analysis (assumes whole paragraph relevant to subject)
        let places = this.nlp(text).places().json();
        for (i in places) {
            this.immediate.push(places[i].text);
        }

        let hyphenateds = this.nlp(text).hyphenated().json();
        for (i in hyphenateds) {
            this.immediate.push(hyphenateds[i].text);
        }

        //get sentences in context
        let tags = this.nlp(text).tag().out('tags');
       
        var pronouns = [subject];
        var tagDoc = [];
        for (var x= 0; x < tags.length; x++) {
            Object.entries(tags[x]).forEach(([k, v]) => {
            //more immediate analysis
                if (v.indexOf("Demonym") > -1) {
                    this.immediate.push(k);
                }
                if (v.indexOf("ProperNoun") > -1) {
                    pronouns.push(k);
                }
               // if (v.indexOf("Adverb") > -1) {
                //    tagDoc.push(k);
               // }
                if (k.length > 5 && v.indexOf("Adjective") > -1) {
                    tagDoc.push(k);
                }
               // if (v.indexOf("Verb") > -1) {
                //    tagDoc.push(k);
                //}
                if (k.length > 5 && v.indexOf("Noun") > -1 && v.indexOf("Singular") > -1) {
                    tagDoc.push(k)
                }
            });
        }

        tagDoc = this.removeRepeats(tagDoc);
        tagDoc = tagDoc.join(" ");
        pronouns = this.removeRepeats(pronouns);
      
        let matches = [];
        sentences.forEach(function (val) {
            val = val.toLowerCase();
            for (p in pronouns) {
                if (val.indexOf(pronouns[p]) > -1) {
                    matches.push(val);
                }
            }
        });
        matches = this.removeRepeats(matches);
      
        var results = [];
        for (m in matches) {
            let bigramDict = this.nlp(matches[m]).bigrams(); 
             for (bi in bigramDict) {
                results.push(this.nlp(tagDoc).lookup(bigramDict[bi].normal).out('array'));
             }
            
        }
        return this.removeRepeats((this.processArrays(results) + "," + this.immediate).split(","));

    },
    processArrays: function(input) {
        let output = [];
        for (i in input) {
            if (input[i].length > 0) {
                output.push(input[i]);
            }
        }
        return output.join(",");
    },
    removeRepeats: function(input) {
        var uniqueList= input.filter(function(item,i,allItems) {
            return i==allItems.indexOf(item);
        });
        return uniqueList;
    },
}
const a = new Analyse();
a.init();
const b = a.sentences("Arachne (/əˈrækniː/; from Ancient Greek: ᾰ̓ρᾰ́χνη, romanized: arákhnē, lit. 'spider', cognate with Latin araneus)[1] is the protagonist of a tale in Roman mythology known primarily from the version told by the Roman poet Ovid (43 BCE–17 CE), which is the earliest extant source for the story.[2] In Book Six of his epic poem Metamorphoses, Ovid recounts how the talented mortal Arachne, daughter of Idmon, challenged Athena, goddess of wisdom and crafts, to a weaving contest. When Athena could find no flaws in the tapestry Arachne had woven for the contest, the goddess became enraged and beat the girl with her shuttle. After Arachne hanged herself out of shame, she was transformed into a spider. The myth both provides an aetiology of spiders' web-spinning abilities and is a cautionary tale warning mortals not to place themselves on an equal level with the gods.");
console.log(b)